FROM golang:1.20 AS base

WORKDIR /app

# Install Air for hot reloading
RUN go install github.com/cosmtrek/air@latest

# Copy Go modules and download dependencies
COPY src/go.mod src/go.sum ./
RUN go mod download

# Copy the application source code
COPY src/ ./

# Final stage for running the application with Air
FROM golang:1.20

WORKDIR /app

# Copy Air binary from the base stage
COPY --from=base /go/bin/air /usr/local/bin/air

# Copy the application source code
COPY --from=base /app /app

# Expose the application port
EXPOSE 8080

# Command to start the application with Air
CMD ["air"]