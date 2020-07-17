FROM alpine:3.12
RUN apk add -U --no-cache ca-certificates

WORKDIR /data
ENV HOME=/data
COPY ./bin/staked /usr/bin/staked
COPY ./bin/stakecli /usr/bin/stakecli
COPY ./docker/entry-point.sh ./entry-point.sh
EXPOSE 26657

ENTRYPOINT [ "/usr/bin/staked" ]
CMD ["start", "--pruning", "nothing"]
