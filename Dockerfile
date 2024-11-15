# Use official Golang image
FROM golang:1.19

# Set working directory
WORKDIR /app

# Copy project files
COPY . .

# Download dependencies
RUN go mod tidy

# Build the Go app
RUN go build -o main .

# Expose port 8080
EXPOSE 8080

# Command to run the application
CMD ["./main"]
