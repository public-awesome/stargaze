FROM alpine:3.15
RUN apk add -U --no-cache ca-certificates

WORKDIR /data
ENV HOME=/data
COPY /stargaze//bin/starsd /usr/bin/starsd
COPY /stargaze/docker/entry-point.sh ./entry-point.sh
EXPOSE 26657

CMD ["starsd", "start", "--pruning", "nothing", "--log_format", "json"]
