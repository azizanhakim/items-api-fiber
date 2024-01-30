# Use an official Golang runtime as a parent image
FROM golang:1.21.5 AS builder

# Set the working directory in the container
WORKDIR /build

# Copy the rest of the application source code
COPY . .

# Build the application
RUN --mount=type=ssh \
    --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build,id=items-api-fiber \
    go mod tidy && CGO_ENABLED=0 GOOS=linux go build -v -installsuffix 'static' -o /build/items-api-fiber .
# RUN CGO_ENABLED=0 go build -o main .

# Use a minimal base image for the final stage
FROM alpine:latest

WORKDIR "/app"

# Install PostgreSQL client libraries
RUN apk add --no-cache postgresql-client

RUN apk --no-cache add curl ca-certificates
RUN apk --no-cache add tzdata
RUN mkdir -p /app/logs

# Set environment variables with default values
ARG DB_HOST=host.docker.internal
ARG DB_NAME=fiber-item
ARG DB_USER=postgres
ARG DB_PASSWORD=azizanhakim
ARG DB_PORT=5432

# Set environment variables
ENV DB_HOST=$DB_HOST
ENV DB_NAME=$DB_NAME
ENV DB_USER=$DB_USER
ENV DB_PASSWORD=$DB_PASSWORD
ENV DB_PORT=$DB_PORT
ENV PORT=$APP_PORT
ENV TZ=Asia/Jakarta

# Copy the built executable from the builder stage
COPY --from=builder /build/.env /app
COPY --from=builder /build/items-api-fiber /app

# Expose the port the app runs on
EXPOSE 8000

# Command to run the application
CMD ["/app/items-api-fiber"]
