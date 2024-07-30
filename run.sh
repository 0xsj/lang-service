#!/bin/sh

directories=("auth")

for dir in "${directories[@]}"; do
    echo "initializing services in $dir"
    cd $dir || { echo "Directory $dir not found"; exit 1; }
    go run src/main.go || { echo "Failed to run server in $dir"; exit 1; }
    cd ..
done

echo "services initialized"
