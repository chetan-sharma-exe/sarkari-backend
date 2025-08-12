FROM golang:1.24-alpine3.21

# Install git & certs
RUN apk add --no-cache git ca-certificates

WORKDIR /app

# Copy go.mod & go.sum first (cache optimization)
COPY go.mod go.sum ./
RUN go mod download

# Copy all code
COPY . .

# Build binary
RUN go build -o main ./cmd

# Expose API port
EXPOSE 8080

# Run app
ENTRYPOINT ["./main"]