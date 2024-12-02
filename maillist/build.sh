
#!/bin/bash

# Load variables from .env file
set -a
source .env
set +a

# Build the Docker image
docker build -f Dockerfile.go.build \
  --build-arg PROTOC_VERSION="$PROTOC_VERSION" \
  --build-arg BUF_VERSION="$BUF_VERSION" \
  --build-arg GO_VERSION="$GO_VERSION" \
  --build-arg PROTOC_GEN_GO_VERSION="$PROTOC_GEN_GO_VERSION" \
  --build-arg PROTOC_GEN_GO_GRPC_VERSION="$PROTOC_GEN_GO_GRPC_VERSION" \
  -t "$IMAGE_NAME" .