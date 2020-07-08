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
RUN go build -o main main.go

# Run the app
CMD ["/build/main"]
