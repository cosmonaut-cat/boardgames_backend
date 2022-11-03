FROM golang:1.18-bullseye AS builder

WORKDIR /go/src/github.com/cosmonaut-cat/boardgames_backend/

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ./boardgames_api github.com/cosmonaut-cat/boardgames_backend/cmd/backend_api

FROM debian:bullseye-slim

RUN apt-get update && apt-get install -y --no-install-recommends ca-certificates

WORKDIR /root/

COPY --from=builder /go/src/github.com/cosmonaut-cat/boardgames_backend/boardgames_api /usr/local/bin/boardgames_api

EXPOSE 3030/tcp

ENTRYPOINT [ "boardgames_api" ]
