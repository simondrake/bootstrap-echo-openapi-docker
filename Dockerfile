FROM golang:alpine

# Set necessary environment variables needed for the image
ENV GO111MODULE=on \
    GCO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Move to working directory /build
WORKDIR /build

# Copy and download dependency using go mod
COPY go.mod .
COPY go.sum .
RUN go mod download

# Copy the code into the container
ADD api ./api
ADD cmd ./cmd
ADD internal ./internal

# Build the application
RUN go build -o bootstrap-echo-openapi-docker cmd/api/main.go

# Move to /app directory as the place for resulting binary folder
WORKDIR /app

# Copy binary from build to main folder
RUN cp /build/bootstrap-echo-openapi-docker .

EXPOSE 8112

# Command to run when starting the container
CMD ["/app/bootstrap-echo-openapi-docker"]
