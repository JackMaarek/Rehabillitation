FROM golang:1.24-alpine

# Enable Go modules
ENV GO111MODULE=on

# Explicit go bin path for binaries
ENV GOBIN=/go/bin
ENV PATH="${GOBIN}:${PATH}"

# Set the working directory
WORKDIR /go/src/github.com/JackMaarek/Rehabillitation/taskAPI

# Copy go.mod and go.sum first
COPY go.mod .
COPY go.sum .

# Download dependencies
RUN go mod download -x

# Install Air (live reload tool)
RUN go install github.com/air-verse/air@latest

# Copy the rest of the application
COPY . .

# Expose the app port
EXPOSE 8000

# Start the hot reload server
CMD ["air"]
