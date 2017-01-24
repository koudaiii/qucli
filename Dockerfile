FROM alpine:3.4

RUN apk add --no-cache --update ca-certificates

COPY bin/qucli /qucli

ENTRYPOINT ["/qucli"]
