# Update the Go version to match your go.mod requirement
FROM golang:1.23.4

# Set the working directory inside the container
WORKDIR /app

# Copy Go module files first
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the application files
COPY . .

# Build the Go application
RUN go build -o app

# Set the command to run the application
CMD ["/app/app"]
