# -----------------------------
# Build container
# -----------------------------
FROM golang:1.14.4 AS builder
WORKDIR /app


COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o image_splitter .


# -----------------------------
# Final container
# -----------------------------
FROM alpine:latest 

COPY --from=builder /app/image_splitter /usr/local/bin/image_splitter