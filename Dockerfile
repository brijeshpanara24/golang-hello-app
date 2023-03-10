FROM golang:1.14 as build

# Add a work directory
WORKDIR /go/src/app

# Cache and install dependencies
COPY . .

# Build app
RUN go build -v -o /app .

# Mention all exposed ports
# EXPOSE PORT1
# EXPOSE PORT2

# Expose port
EXPOSE 8080

# Now copy it into our base image.
FROM gcr.io/distroless/base
COPY --from=build /app /app
CMD ["/app"]