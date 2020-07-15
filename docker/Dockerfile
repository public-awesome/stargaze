FROM alpine:3.12
RUN apk add -U --no-cache ca-certificates

WORKDIR /root
COPY ./bin/staked /usr/bin/staked
COPY ./bin/stakecli /usr/bin/stakecli
COPY ./docker/entry-point.sh ./entry-point.sh
EXPOSE 26657

ENTRYPOINT [ "./entry-point.sh" ]
