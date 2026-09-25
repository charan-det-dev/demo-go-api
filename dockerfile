# Use an official Golang runtime as a parent image with the specific version
FROM golang:1.24.0-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files
COPY go.mod go.sum ./

# Download Go module dependencies
RUN go mod download -x

# Copy the Go source files
COPY . .

# Build the Go application
RUN go build -o main .

# Create a minimal runtime image
FROM alpine:latest 

# Set the working directory in the runtime image
WORKDIR /app

# Copy the built executable from the builder stage
COPY --from=builder /app/main .

# Expose the port that your application listens on
EXPOSE 8080

# Command to run the executable
CMD ["./main"]