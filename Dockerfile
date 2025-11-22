FROM golang:1.25-alpine

WORKDIR /app

# Copy source code
COPY . .

# Build the application
RUN go build -o parking-app ./cmd/parking-cli

# Expose port
EXPOSE 8080

# Run the application
CMD ["./parking-app"]