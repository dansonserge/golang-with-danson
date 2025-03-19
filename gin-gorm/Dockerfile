# Use the official Golang image as the base image
FROM golang:1.24-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the Go modules files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code to the container
COPY . .

# Expose the application port
EXPOSE 8080

# Build the application
RUN go build -o gin-gorm

# Run the application
CMD ["./gin-gorm"]
