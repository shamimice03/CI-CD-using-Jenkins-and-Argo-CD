# Use an official Golang runtime as a parent image
FROM golang:alpine AS builder   

# Set the working directory to /go/src/app
WORKDIR /go/src/app

# Copy go.mod and go.sum files into the container
#COPY go.mod .
#COPY go.sum .

# Copy the current directory contents into the container at /go/src/app
COPY . .

# Download and install application dependencies
RUN go mod download


# Build the Golang binary
RUN go build -o main .

# Expose port 8080 for the HTTP server to listen on
EXPOSE 8080

# Run the binary when the container starts
# CMD ["./main"]

# new stage -------------------
FROM alpine
# Copy artifact from builder stage                   
COPY --from=builder /main .   
CMD ["./main"]
