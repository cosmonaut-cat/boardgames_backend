FROM golang:1.18-bullseye AS builder

WORKDIR /go/src/github.com/cosmonaut-cat/boardgames_backend/

COPY ./ ./

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ./front_api github.com/cosmonaut-cat/boardgames_backend/cmd/front_api

FROM debian:bullseye-slim

RUN apt-get update && apt-get install -y --no-install-recommends ca-certificates

WORKDIR /root/

COPY --from=builder /go/src/github.com/cosmonaut-cat/boardgames_backend/front_api /usr/local/bin/front_api

EXPOSE 3031/tcp

ENTRYPOINT [ "front_api" ]
