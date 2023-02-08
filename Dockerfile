# Use an official Go image as the base image
FROM golang:1.19

# Set the working directory in the container
WORKDIR /app

# Copy the source code to the container
COPY . .

# Build the Go application
RUN go build -o bootstrap-cli .

# Specify the command to run the application when the container starts
CMD ["./bootstrap-cli"]