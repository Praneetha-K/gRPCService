# gRPC User Service App

This a simple Go application demonstrating a gRPC service.

## Table of Contents

- [Prerequisites](#prerequisites)
- [Getting Started](#getting-started)
- [Usage](#usage)
- [Docker](#docker)

## Prerequisites

- Go (1.16 or higher)
- Protobuf compiler (protoc)
- [Docker](https://www.docker.com/)

## Getting Started

1. Clone the repository:

    ```bash
    git clone https://github.com/Praneetha-K/gRPCService.git
    cd GRPC_Totality_Corp
    ```

2. Generate the Go files from the proto file:

    ```bash
    protoc --go_out=. --go-grpc_out=. user.proto
    ```

3. Build and run the application:

    ```bash
    go run main.go
    ```

The gRPC service should now be running on `localhost:8080`.

## Usage

To use the service, you can create a gRPC client
    ```bash
    grpcurl -plaintext -d '{"user_id": 1}' localhost:8080 UserService/GetUserDetails
    ```

## Docker
```bash
    docker build -t grpc-user-service .
    docker run -p 8080:8080 grpc-user-service
```