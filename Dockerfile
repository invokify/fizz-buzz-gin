############################
# STEP 1 build executable binary
############################
FROM golang:1.20.1-alpine AS builder

# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git

# Create appuser.
RUN adduser -D -g '' appuser

# Create data folder.
RUN mkdir /data

# Create data folder.
RUN mkdir /providers

# Copy files
WORKDIR /app
COPY . .

# Using go mod.
RUN go mod verify

# Build the binary.
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-w -s" -o fizzBuzzApp ./cmd/api

############################
# STEP 2 build a small image
############################
# FROM scratch
FROM builder
# Import the user and group files from the builder.
COPY --from=builder /etc/passwd /etc/passwd

# Copy our static executable.
COPY --from=builder /app/fizzBuzzApp /app

# Use an unprivileged user.
USER appuser

# Port on which the service will be exposed.
EXPOSE 80

# Run the product-service binary.
ENTRYPOINT ["/app/fizzBuzzApp"]