# Stage 1
FROM golang:alpine as builder

WORKDIR /app

# Copy only the go.mod and go.sum files to download dependencies
COPY go.mod .
COPY go.sum .

# Download dependencies
RUN go mod download

# Copy the rest of the application source code
COPY . .

# Build the application
RUN go build -o main .

# Stage 2
FROM alpine
RUN adduser -S -D -H -h /app appuser
USER appuser
WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/main .

EXPOSE 3000
CMD ["./main"]
