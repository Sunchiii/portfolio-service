# Start from a base Go image
FROM golang:1.20-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files
COPY go.mod go.sum ./

# Download the Go module dependencies
RUN go mod download
RUN go mod tidy

# Copy the rest of the application source code
COPY . .

# Build the Go application
RUN go build ./api/main.go

# Expose the port that the API will listen on
EXPOSE 8080

# Set the command to run the API binary
CMD ["./main"]
