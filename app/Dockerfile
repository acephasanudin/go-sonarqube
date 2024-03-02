# Use the official Golang image from Docker Hub
FROM golang:1.22.0

# Set the working directory inside the container
WORKDIR /usr/src/app

# Create a non-root user
RUN useradd -u 10001 gouser

# Copy the local code into the container
COPY ./go.mod .
COPY ./main.go .
COPY ./main_test.go .

# Change ownership to the non-root user
RUN chown -R gouser:gouser /usr/src/app

# Switch to the non-root user
USER myuser

# Build the Go application
RUN go build -o app

# Expose a port (optional, depends on your application)
EXPOSE 8080

# Command to run the executable
CMD ["./app"]
