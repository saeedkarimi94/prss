# Use an official Golang runtime as a base image
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the local package files to the container's workspace
COPY . .

# Download and install any dependencies your application needs
RUN go mod download

# Build the Go application
RUN go build -o myapp

# Expose the port the application will run on
EXPOSE 8080

# Define the command to run your application
CMD ["./myapp"]
