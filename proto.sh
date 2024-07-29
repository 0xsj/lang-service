#!/bin/sh

# Define arrays for Go services and the gateway
goServices=("auth")
gateway=("auth" "user" "notification" "identity")

# Cleanup function to remove old proto files and create a new directory
cleanup() {
  folderPath="$1/proto"
  echo "Cleaning up $folderPath"
  if [ -d "$folderPath" ]; then
    rm -r "$folderPath"
  fi
  mkdir -p "$folderPath"
}

# Generate Go files from .proto files
generateGo() {
    path="$1/proto"
    protoFile="$2.proto"
    echo "Generating Go files for $protoFile in $path"
    protoc --go_out="$path" --go-grpc_out="$path" --proto_path=proto "$protoFile"
}

# Generate Go files for each service
for service in "${goServices[@]}"; do
    echo "Processing service: $service"
    cleanup "$service"
    generateGo "$service" "$service"
done

# Generate Go files for the gateway
echo "Processing conduit"
cleanup "conduit"
for file in "${gateway[@]}"; do
    generateGo "conduit" "$file"
done
