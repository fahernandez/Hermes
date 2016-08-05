FROM alpine:3.4

RUN apk add --update ca-certificates && rm -rf /var/cache/apk/*

COPY hermes /hermes

ENTRYPOINT ["/hermes"]

EXPOSE 8001
