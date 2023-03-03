FROM golang:1.16.5

# Add a work directory
WORKDIR /app

# Cache and install dependencies
COPY . .
RUN go mod download

# Build app
RUN go build -o app

# Mention all exposed ports
# EXPOSE PORT1
# EXPOSE PORT2

# Expose port
EXPOSE 8080

# Exec built binary
CMD ./app