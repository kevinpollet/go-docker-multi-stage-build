FROM golang:1.13 AS builder
WORKDIR /build
COPY go.* ./
RUN go mod download
COPY . .
RUN go build -o app

FROM gcr.io/distroless/base
COPY --from=builder /build/app .
ENTRYPOINT [ "./app" ]
