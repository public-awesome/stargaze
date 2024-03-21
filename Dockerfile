# docker build . -t publicawesome/stargaze:latest
# docker run --rm -it publicawesome/stargaze:latest /bin/sh
FROM golang:1.22.1-alpine3.19 AS go-builder


RUN set -eux; apk add --no-cache ca-certificates build-base git;

# TARGETPLATFORM build argument provided by buildx: e.g., linux/amd64, linux/arm64. etc.
ARG TARGETPLATFORM=linux/amd64

# NOTE: add these to run with LEDGER_ENABLED=true
# RUN apk add libusb-dev linux-headers

WORKDIR /code
COPY . /code/

# See https://github.com/CosmWasm/wasmvm/releases
# Download the correct version of libwasmvm for the given platform and verify checksum
RUN case "${TARGETPLATFORM}" in \
      "linux/amd64") \
        WASMVM_URL="https://github.com/CosmWasm/wasmvm/releases/download/v1.5.2/libwasmvm_muslc.x86_64.a" && \
        WASMVM_CHECKSUM="e660a38efb2930b34ee6f6b0bb12730adccb040b6ab701b8f82f34453a426ae7" \
        ;; \
      "linux/arm64") \
        WASMVM_URL="https://github.com/CosmWasm/wasmvm/releases/download/v1.5.2/libwasmvm_muslc.aarch64.a" && \
        WASMVM_CHECKSUM="e78b224c15964817a3b75a40e59882b4d0e06fd055b39514d61646689cef8c6e" \
        ;; \
      *) \
        echo "Unsupported platform: ${TARGETPLATFORM}" ; \
        exit 1 \
        ;; \
    esac && \
    wget "${WASMVM_URL}" -O /lib/libwasmvm_muslc.a && \
    echo "${WASMVM_CHECKSUM}  /lib/libwasmvm_muslc.a" | sha256sum -c

# force it to use static lib (from above) not standard libgo_cosmwasm.so file
RUN  LEDGER_ENABLED=false BUILD_TAGS=muslc LINK_STATICALLY=true  make build


# --------------------------------------------------------
FROM alpine:3.19

COPY --from=go-builder /code/bin/starsd /usr/bin/starsd
RUN apk add -U --no-cache ca-certificates
WORKDIR /data
ENV HOME=/data
COPY ./docker/entry-point.sh ./entry-point.sh
# rest server
EXPOSE 1317
# tendermint p2p
EXPOSE 26656
# tendermint rpc
EXPOSE 26657


CMD ["starsd", "start", "--pruning", "nothing", "--log_format", "json"]
