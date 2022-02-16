package wasm

import "fmt"

type ClaimAction string

const (
	ClaimActionMintNFT = "mint_nft"
	ClaimActionBidNFT  = "bid_nft"
)

type ClaimFor struct {
	Address string      `json:"address"`
	Action  ClaimAction `json:"action"`
}

func (a ClaimAction) Valid() error {
	if a == ClaimActionMintNFT || a == ClaimActionBidNFT {
		return nil
	}
	return fmt.Errorf("invalid action")
}

type ClaimMsg struct {
	ClaimFor *ClaimFor `json:"claim_for,omitempty"`
}
