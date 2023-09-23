# Stage 1: Build the Go application
FROM golang:1.20-alpine AS builder

WORKDIR /app

# set ENV
ENV sd=value

# Copy the Go application source code and Makefile
COPY . .

# Build the Go application with the Makefile
RUN apk update && apk add make
RUN apk update && apk add curl

# Install Swag
RUN go install github.com/swaggo/swag/cmd/swag@latest

# Download Golang-Migrate using curl and install it
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.14.1/migrate.linux-amd64.tar.gz | tar xvz
RUN mv migrate.linux-amd64 /usr/bin/migrate

# Run migrations using Golang-Migrate
RUN make migrate-dev-up

# Build the main application
RUN make app-build


# Stage 2: Create the final lightweight container
# FROM alpine:3.14


# WORKDIR /app
# RUN ls -al

# # Copy the binary from the builder stage
# COPY --from=builder /app/main .
# COPY .env .
# COPY external/migration ./external/migration

# Expose the port the application will listen on
EXPOSE 8801

# Define the command to run the application
CMD ["/app/bin/main"]