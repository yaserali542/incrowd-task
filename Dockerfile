# Start from golang base image
FROM golang:alpine

# Add Maintainer info
LABEL maintainer="Agus Wibawantara"

# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git && apk add --no-cach bash && apk add build-base

# Setup folders
RUN mkdir /app
WORKDIR /app
COPY . .
# Download all the dependencies
RUN go get -d -v ./...

# Install the package
RUN go install -v ./...

# Build the Go app
RUN go build -o /build
EXPOSE 8080

# Run the executable
CMD [ "/build" ]