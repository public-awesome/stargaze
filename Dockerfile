# docker build . -t publicawesome/stargaze:latest
# docker run --rm -it publicawesome/stargaze:latest /bin/sh
FROM golang:1.21.0-alpine3.17 AS go-builder


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
        WASMVM_URL="https://github.com/CosmWasm/wasmvm/releases/download/v1.3.0/libwasmvm_muslc.x86_64.a" && \
        WASMVM_CHECKSUM="b1610f9c8ad8bdebf5b8f819f71d238466f83521c74a2deb799078932e862722" \
        ;; \
      "linux/arm64") \
        WASMVM_URL="https://github.com/CosmWasm/wasmvm/releases/download/v1.3.0/libwasmvm_muslc.aarch64.a" && \
        WASMVM_CHECKSUM="b1610f9c8ad8bdebf5b8f819f71d238466f83521c74a2deb799078932e862722" \
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
FROM alpine:3.17

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
