FROM geekidea/alpine-a:3.9

RUN apk update \
        && apk upgrade \
        && apk add --no-cache \
        ca-certificates \
        && update-ca-certificates 2>/dev/null || true


RUN mkdir /app
## We copy everything in the root directory
## into our /app directory
## We specify that we now wish to execute
## any further commands inside our /ap€p
## directory
WORKDIR /app
## we run go build to compile the binary
## executable of our Go program
COPY example-project /app/main
COPY config.yaml /app/config.yaml
## Our start command which kicks off
## our newly created binary executable
CMD ["/app/main","api","--config","/app/config.yaml"]