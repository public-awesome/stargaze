#!/usr/bin/env bash
#
# scripts/test-v18-patch-fork.sh
#
# One-off local verification that the v18-patch hard fork actually fires on a
# running stargaze-1 chain and rewrites gov params end-to-end. Delete this
# file once you're satisfied — it's a throwaway smoke test, not part of CI.
#
# How it works:
#   - Reads the production v18_patch.UpgradeHeight directly from constants.go
#     (no source patching).
#   - Starts a single-validator stargaze-1 chain with cometbft's genesis
#     `initial_height` set to UpgradeHeight - 15, so the chain produces ~15
#     blocks and then hits the real fork height naturally.
#   - Seeds the broken pre-fork state in genesis (MinDepositRatio = "").
#   - Waits for the chain to cross UpgradeHeight, then asserts the gov
#     params match what RunForkLogic should write.
#
# Requires: go, jq, sed.

set -euo pipefail

CHAIN_ID="stargaze-1"
BLOCKS_BEFORE_FORK=15
DENOM="ustars"
REPO_ROOT="$(cd "$(dirname "$0")/.." && pwd)"
STARGAZE_HOME="$REPO_ROOT/.v18patch-fork-test"
STARSD_BIN="$REPO_ROOT/bin/starsd-v18patch-fork-test"
CONSTANTS_FILE="$REPO_ROOT/app/upgrades/mainnet/v18_patch/constants.go"
LOG_FILE="$REPO_ROOT/.v18patch-fork-test.log"
PID_FILE="$STARGAZE_HOME/starsd.pid"

color_red()   { printf "\033[31m%s\033[0m\n" "$*"; }
color_green() { printf "\033[32m%s\033[0m\n" "$*"; }
log() { printf ">> %s\n" "$*"; }

cleanup() {
    local rc=$?
    log "cleanup (exit $rc)"
    if [[ -f "$PID_FILE" ]]; then
        local pid; pid=$(cat "$PID_FILE" 2>/dev/null || true)
        if [[ -n "${pid:-}" ]] && kill -0 "$pid" 2>/dev/null; then
            kill "$pid" 2>/dev/null || true
            sleep 1
            kill -9 "$pid" 2>/dev/null || true
        fi
        rm -f "$PID_FILE"
    fi
    if [[ $rc -ne 0 && -f "$LOG_FILE" ]]; then
        echo "----- last 40 lines of $LOG_FILE -----"
        tail -n 40 "$LOG_FILE" || true
    fi
}
trap cleanup EXIT

if ! command -v jq >/dev/null; then
    color_red "jq is required"; exit 1
fi

# --- 1. read real UpgradeHeight from source ---------------------------------
# Pull the production value out of constants.go so the test stays in sync
# with whatever the binary will actually use on mainnet.
FORK_HEIGHT=$(grep -E 'UpgradeHeight = int64\([0-9_]+\)' "$CONSTANTS_FILE" \
    | sed -E 's/.*int64\(([0-9_]+)\).*/\1/' | tr -d '_')
if [[ ! "$FORK_HEIGHT" =~ ^[0-9]+$ ]] || (( FORK_HEIGHT < 100 )); then
    color_red "could not read a usable UpgradeHeight from $CONSTANTS_FILE (got '$FORK_HEIGHT')"
    exit 1
fi
INITIAL_HEIGHT=$(( FORK_HEIGHT - BLOCKS_BEFORE_FORK ))
log "production UpgradeHeight = $FORK_HEIGHT"
log "chain will start at height $INITIAL_HEIGHT (~$BLOCKS_BEFORE_FORK blocks before fork)"

# --- 2. build the binary (uses real UpgradeHeight, no source mutation) ------
log "building $STARSD_BIN"
mkdir -p "$REPO_ROOT/bin"
(cd "$REPO_ROOT" && go build -o "$STARSD_BIN" ./cmd/starsd)

# --- 3. init chain -----------------------------------------------------------
log "init chain (chain-id=$CHAIN_ID, home=$STARGAZE_HOME)"
rm -rf "$STARGAZE_HOME"
mkdir -p "$STARGAZE_HOME"

"$STARSD_BIN" config set client chain-id "$CHAIN_ID" --home "$STARGAZE_HOME" >/dev/null
"$STARSD_BIN" config set client keyring-backend test --home "$STARGAZE_HOME" >/dev/null
"$STARSD_BIN" config set client output json --home "$STARGAZE_HOME" >/dev/null

"$STARSD_BIN" keys add validator --keyring-backend test --home "$STARGAZE_HOME" </dev/null >/dev/null 2>&1
VALIDATOR=$("$STARSD_BIN" keys show validator -a --keyring-backend test --home "$STARGAZE_HOME")

"$STARSD_BIN" init test-node --chain-id "$CHAIN_ID" --home "$STARGAZE_HOME" >/dev/null 2>&1

# fast blocks
CONFIG="$STARGAZE_HOME/config/config.toml"
if [[ "$(uname)" == "Linux" ]]; then
    sed -i 's/^timeout_commit = .*/timeout_commit = "500ms"/' "$CONFIG"
else
    sed -i '' 's/^timeout_commit = .*/timeout_commit = "500ms"/' "$CONFIG"
fi

# Patch genesis: bump initial_height past most of the gap to UpgradeHeight,
# install the broken pre-fork MinDepositRatio, and align denoms.
GENESIS="$STARGAZE_HOME/config/genesis.json"
log "patching genesis (initial_height=$INITIAL_HEIGHT, denom=$DENOM, MinDepositRatio=\"\")"
contents=$(jq \
    --arg denom "$DENOM" \
    --arg ih "$INITIAL_HEIGHT" '
    .initial_height = $ih
    | .app_state.staking.params.bond_denom = $denom
    | .app_state.crisis.constant_fee.denom = $denom
    | .app_state.mint.params.mint_denom = $denom
    | .app_state.gov.params.min_deposit[0].denom = $denom
    | .app_state.gov.params.expedited_min_deposit[0].denom = $denom
    | .app_state.gov.params.voting_period = "60s"
    | .app_state.gov.params.expedited_voting_period = "30s"
    | .app_state.gov.params.min_deposit_ratio = ""
' "$GENESIS")
echo "$contents" > "$GENESIS"

"$STARSD_BIN" genesis add-genesis-account "$VALIDATOR" "10000000000000$DENOM" --home "$STARGAZE_HOME" >/dev/null
"$STARSD_BIN" genesis gentx validator "10000000000$DENOM" --chain-id "$CHAIN_ID" --keyring-backend test --home "$STARGAZE_HOME" >/dev/null 2>&1
"$STARSD_BIN" genesis collect-gentxs --home "$STARGAZE_HOME" >/dev/null 2>&1
"$STARSD_BIN" genesis validate-genesis --home "$STARGAZE_HOME" >/dev/null

# --- 4. start chain ---------------------------------------------------------
log "starting chain (logs: $LOG_FILE)"
"$STARSD_BIN" start --home "$STARGAZE_HOME" > "$LOG_FILE" 2>&1 &
CHAIN_PID=$!
echo "$CHAIN_PID" > "$PID_FILE"

# wait for first block (at INITIAL_HEIGHT)
log "waiting for the chain to produce its first block at height >= $INITIAL_HEIGHT"
h=0
for _ in $(seq 1 60); do
    if ! kill -0 "$CHAIN_PID" 2>/dev/null; then
        color_red "chain process $CHAIN_PID died unexpectedly"
        exit 1
    fi
    h=$("$STARSD_BIN" status --home "$STARGAZE_HOME" 2>/dev/null \
        | jq -r '(.sync_info // .SyncInfo).latest_block_height // "0"' 2>/dev/null || echo "0")
    if [[ "$h" =~ ^[0-9]+$ ]] && (( h >= INITIAL_HEIGHT )); then break; fi
    sleep 1
done
log "current height: $h"
if [[ ! "$h" =~ ^[0-9]+$ ]] || (( h < INITIAL_HEIGHT )); then
    color_red "chain never reached its initial_height; check $LOG_FILE"
    exit 1
fi

# --- 5. assert pre-fork MinDepositRatio is empty/missing --------------------
log "asserting pre-fork MinDepositRatio is empty/missing"
PRE=$("$STARSD_BIN" query gov params --home "$STARGAZE_HOME" --output json | jq -r '.params.min_deposit_ratio')
log "  pre-fork MinDepositRatio = '$PRE'"
if [[ -n "$PRE" && "$PRE" != "null" ]]; then
    color_red "FAIL: pre-fork MinDepositRatio expected to be empty/null, got '$PRE'"
    exit 1
fi

# --- 6. wait past UpgradeHeight ---------------------------------------------
log "waiting for height > $FORK_HEIGHT"
for _ in $(seq 1 120); do
    if ! kill -0 "$CHAIN_PID" 2>/dev/null; then
        color_red "chain process died while waiting for fork height"
        exit 1
    fi
    h=$("$STARSD_BIN" status --home "$STARGAZE_HOME" 2>/dev/null \
        | jq -r '(.sync_info // .SyncInfo).latest_block_height // "0"' 2>/dev/null || echo "0")
    if [[ "$h" =~ ^[0-9]+$ ]] && (( h > FORK_HEIGHT )); then break; fi
    sleep 1
done
log "  reached height $h"

# --- 7. assert post-fork MinDepositRatio == "0.010000000000000000" -----------
log "asserting post-fork gov params"
PARAMS_JSON=$("$STARSD_BIN" query gov params --home "$STARGAZE_HOME" --output json)
POST_RATIO=$(echo "$PARAMS_JSON" | jq -r '.params.min_deposit_ratio')
POST_MIN_DEPOSIT=$(echo "$PARAMS_JSON" | jq -r '.params.min_deposit[0].amount + .params.min_deposit[0].denom')
POST_QUORUM=$(echo "$PARAMS_JSON" | jq -r '.params.quorum')

log "  MinDepositRatio  = '$POST_RATIO'"
log "  MinDeposit       = '$POST_MIN_DEPOSIT'"
log "  Quorum           = '$POST_QUORUM'"

failed=0
if [[ "$POST_RATIO" != "0.010000000000000000" ]]; then
    color_red "FAIL: post-fork MinDepositRatio expected 0.010000000000000000, got '$POST_RATIO'"
    failed=1
fi
if [[ "$POST_MIN_DEPOSIT" != "500000000000ustars" ]]; then
    color_red "FAIL: post-fork MinDeposit expected 500000000000ustars, got '$POST_MIN_DEPOSIT'"
    failed=1
fi
if [[ "$POST_QUORUM" != "0.200000000000000000" ]]; then
    color_red "FAIL: post-fork Quorum expected 0.200000000000000000, got '$POST_QUORUM'"
    failed=1
fi

if [[ $failed -eq 0 ]]; then
    color_green ""
    color_green "=== PASS: v18-patch fork applied at height $FORK_HEIGHT on $CHAIN_ID ==="
    color_green ""
    exit 0
fi
exit 1
