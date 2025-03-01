# Use the official Go image based on Alpine Linux
FROM golang:1.24-alpine

# Set the working directory inside the container
WORKDIR /app

# Install Air for live reloading
RUN go install github.com/air-verse/air@latest

# Copy go.mod and go.sum to leverage Docker layer caching
COPY go.* ./

# Install Git (needed to fetch dependencies) and download modules
RUN apk add --no-cache git
RUN go mod download

# Copy the entire project into the container
COPY . .

# Expose the port your application runs on
EXPOSE 8080

# Run Air with the specified configuration file
CMD ["air", "-c", ".air.toml"]
