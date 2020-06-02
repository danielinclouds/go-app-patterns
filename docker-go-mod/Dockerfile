# Dockerfile for go applications using go modules

# -----------------------------
# Build container
# -----------------------------
FROM golang:1.14.4 AS builder
WORKDIR /app

# Download dependencies
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .


# -----------------------------
# Final container
# -----------------------------
FROM alpine:latest 

ARG PORT
ENV APP_PORT=${PORT}

RUN apk --no-cache add ca-certificates

USER 65534
WORKDIR /app/
COPY --from=builder /app/app .
EXPOSE ${PORT}
CMD ["./app"]
