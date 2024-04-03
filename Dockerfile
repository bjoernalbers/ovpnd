FROM golang:1.22 AS builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 go build -o /ovpnd
FROM scratch
COPY --from=builder /ovpnd /ovpnd
ENTRYPOINT ["/ovpnd"]
