# syntax=docker/dockerfile:1
FROM golang:1.16 AS builder
WORKDIR /go/src/github.com/nguli-team/bakalo/
COPY . .
RUN go get -d -v .
RUN make build

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /app/
COPY --from=builder /go/src/github.com/nguli-team/bakalo/bin .
CMD ["./bakalo", "serve"]
