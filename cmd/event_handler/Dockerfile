FROM golang:1.18-bullseye AS builder

WORKDIR /go/src/github.com/cosmonaut-cat/boardgames_backend/

COPY ./ ./

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ./event_handler github.com/cosmonaut-cat/boardgames_backend/cmd/event_handler

FROM debian:bullseye-slim

RUN apt-get update && apt-get install -y --no-install-recommends ca-certificates

WORKDIR /root/

COPY --from=builder /go/src/github.com/cosmonaut-cat/boardgames_backend/event_handler /usr/local/bin/event_handler

EXPOSE 3030/tcp

ENTRYPOINT [ "event_handler" ]
