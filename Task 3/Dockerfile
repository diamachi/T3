# Start from the official Go image
FROM golang:1.20-alpine

# Set working directory
WORKDIR /app

# Copy go.mod and source code
COPY go.mod ./
COPY main.go ./

# Build the binary
RUN go build -o main .

# Expose port
EXPOSE 8080

# Run the app
CMD ["./main"]
