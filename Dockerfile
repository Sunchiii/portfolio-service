# Start from a base Go image
FROM golang:1.20-alpine

# Set the working directory inside the container
WORKDIR /app

# set env variable
RUN export PG_PORT=5432
RUN export PG_USER=postgres
RUN export PG_PASSWORD=postgrespw
RUN export PG_DATABASE=portfoliodb
RUN export PG_HOST=0.0.0.0
RUN export PORT=8080


RUN export SIGNATURE_KEY=sonfolioSignatureBySunchiii@github.com
RUN export TOKEN_EXP=24

# Copy the Go module files
COPY go.mod go.sum ./

# Download the Go module dependencies
RUN go mod download

# Copy the rest of the application source code
COPY . .

# Build the Go application
RUN go build -o ./server ./api/

# Expose the port that the API will listen on
EXPOSE 8080

# Set the command to run the API binary
CMD ["./server"]
