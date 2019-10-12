FROM golang:1.13-alpine AS builder
WORKDIR build/src
COPY . .
RUN go build -o /build/app .

FROM gcr.io/distroless/static
WORKDIR app
COPY --from=builder /build/app .
ENTRYPOINT [ "/app" ]
