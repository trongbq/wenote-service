FROM golang:1.14.4-alpine

# Move to working directory
WORKDIR /build

# Copy and download dependency
COPY go.mod .
COPY go.sum .
RUN go mod download

# Copy the code into the container
COPY . .

# Build the app
RUN go build -tags api -o ./bin/wenote-api ./cmd/wenote-api/.

# Run the app
CMD ["/build/bin/wenote-api"]
