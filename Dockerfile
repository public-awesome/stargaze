# docker build . -t publicawesome/stargaze:latest
# docker run --rm -it publicawesome/stargaze:latest /bin/sh
FROM golang:1.22.7-alpine3.19 AS go-builder


RUN set -eux; apk add --no-cache ca-certificates build-base git;

# TARGETPLATFORM build argument provided by buildx: e.g., linux/amd64, linux/arm64. etc.
ARG TARGETPLATFORM=linux/amd64

# NOTE: add these to run with LEDGER_ENABLED=true
# RUN apk add libusb-dev linux-headers

WORKDIR /code
COPY . /code/

# See https://github.com/CosmWasm/wasmvm/releases
# Download the correct version of libwasmvm for the given platform and verify checksum
ADD https://github.com/CosmWasm/wasmvm/releases/download/v2.2.1/libwasmvm_muslc.x86_64.a /lib/libwasmvm_muslc.x86_64.a
ADD https://github.com/CosmWasm/wasmvm/releases/download/v2.2.1/libwasmvm_muslc.aarch64.a /lib/libwasmvm_muslc.aarch64.a
RUN echo "a4a3d09b36fabb65b119d5ba23442c23694401fcbee4451fe6b7e22e325a4bac /lib/libwasmvm_muslc.x86_64.a" | sha256sum -c
RUN echo "faea4e15390e046d2ca8441c21a88dba56f9a0363f92c5d94015df0ac6da1f2d /lib/libwasmvm_muslc.aarch64.a" | sha256sum -c

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
