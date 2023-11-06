# STAGE 1: building the executable
FROM golang:1.21 AS builder

ENV GOPATH /go 
ENV GO111MODULE on
ENV CGO_ENABLED 0

COPY . /go/src/github.com/phbpx/mockit
WORKDIR /go/src/github.com/phbpx/mockit

RUN go mod download
RUN go build -installsuffix 'static' -o /go/bin/mockit ./cmd/main.go

# STAGE 2: build the container to run
FROM gcr.io/distroless/static

LABEL maintainer="Paulo Bortolotti <bortolotti.paulo@gmail.com>"
USER nonroot:nonroot

COPY --from=builder --chown=nonroot:nonroot /go/bin/mockit /mockit

ENTRYPOINT ["/mockit"]