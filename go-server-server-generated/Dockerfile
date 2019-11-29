# Start from the latest golang base image
FROM golang:latest

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY ./ .

# Add package mux
RUN go get github.com/BurntSushi/toml
RUN go get -v -u github.com/gorilla/mux
RUN go get gopkg.in/mgo.v2

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN go build -o main .

EXPOSE 8082