# # Start with a lightweight Go image
# FROM golang:1.23 AS builder

# WORKDIR /app

# # Copy Go modules and install dependencies
# COPY go.mod go.sum ./
# RUN go mod download

# # Copy the source code
# COPY . .

# # Build the Go app, specifying the output directory
# RUN go build -o /app/multiplayer-mode-usage

# # Run stage
# FROM alpine:latest

# WORKDIR /root/

# # Copy the binary from the builder stage
# COPY --from=builder /app/multiplayer-mode-usage .

# # Ensure the binary has execute permissions
# RUN chmod +x /root/multiplayer-mode-usage

# # Expose port 8080
# EXPOSE 8080

# # Start the application
# # CMD ["./multiplayer-mode-usage"]

# syntax=docker/dockerfile:1

FROM golang:1.23

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code. Note the slash at the end, as explained in
# https://docs.docker.com/reference/dockerfile/#copy

COPY . .

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /multiplayer-mode-usage

# Optional:
# To bind to a TCP port, runtime parameters must be supplied to the docker command.
# But we can document in the Dockerfile what ports
# the application is going to listen on by default.
# https://docs.docker.com/reference/dockerfile/#expose
EXPOSE 8080

# Run
CMD ["/multiplayer-mode-usage"]