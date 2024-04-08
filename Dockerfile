FROM golang:1.22 AS builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 make
FROM scratch
COPY --from=builder /app/ovpnd /ovpnd
EXPOSE 80
EXPOSE 443
ENTRYPOINT ["/ovpnd"]
