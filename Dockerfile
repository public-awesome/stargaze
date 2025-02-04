# docker build . -t publicawesome/stargaze:latest
# docker run --rm -it publicawesome/stargaze:latest /bin/sh
FROM golang:1.23.5-alpine3.20 AS go-builder


RUN set -eux; apk add --no-cache ca-certificates build-base git;

# TARGETPLATFORM build argument provided by buildx: e.g., linux/amd64, linux/arm64. etc.
ARG TARGETPLATFORM=linux/amd64

# NOTE: add these to run with LEDGER_ENABLED=true
# RUN apk add libusb-dev linux-headers

WORKDIR /code
COPY . /code/

# See https://github.com/CosmWasm/wasmvm/releases
# Download the correct version of libwasmvm for the given platform and verify checksum
ADD https://github.com/CosmWasm/wasmvm/releases/download/v2.1.5/libwasmvm_muslc.x86_64.a /lib/libwasmvm_muslc.x86_64.a
ADD https://github.com/CosmWasm/wasmvm/releases/download/v2.1.5/libwasmvm_muslc.aarch64.a /lib/libwasmvm_muslc.aarch64.a
RUN echo "c6612d17d82b0997696f1076f6d894e339241482570b9142f29b0d8f21b280bf /lib/libwasmvm_muslc.x86_64.a" | sha256sum -c
RUN echo "1bad0e3f9b72603082b8e48307c7a319df64ca9e26976ffc7a3c317a08fe4b1a /lib/libwasmvm_muslc.aarch64.a" | sha256sum -c

# force it to use static lib (from above) not standard libgo_cosmwasm.so file
RUN  LEDGER_ENABLED=false BUILD_TAGS=muslc LINK_STATICALLY=true  make build


# --------------------------------------------------------
FROM alpine:3.20

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
