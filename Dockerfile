FROM golang:1.13 AS builder
WORKDIR /go/src/app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o /go/bin/app

FROM gcr.io/distroless/base
COPY --chown=65534:65534 --from=builder /go/bin/app .
USER 65534
ENTRYPOINT [ "./app" ]
