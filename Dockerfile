FROM golang:1.22 AS builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 make build
FROM scratch
COPY --from=builder /app/ovpnd /ovpnd
ENTRYPOINT ["/ovpnd"]
