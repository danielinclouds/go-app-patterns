# -----------------------------
# Build container
# -----------------------------
FROM golang:1.14.4 AS builder
WORKDIR /app


COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o printer .


# -----------------------------
# Final container
# -----------------------------
FROM alpine:latest 

COPY --from=builder /app/printer /usr/local/bin/printer