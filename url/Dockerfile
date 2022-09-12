# STAGE: BUILD
FROM golang:1.18 as builder

WORKDIR /app

# Source folders
COPY internal url/internal
COPY cmd url/cmd

# Files in the root, therefore copy them from the root before building the image
# ps: if you using the makefile, it will do it for you
COPY go.mod .
COPY go.sum .
COPY Makefile .

RUN make build_url

# STAGE: TARGET
FROM alpine:latest

WORKDIR /app
COPY --from=builder /app/bin/url /app/url
COPY migration /app/migration

ENTRYPOINT ["/app/url"]
