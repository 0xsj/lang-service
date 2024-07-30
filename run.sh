#!/bin/sh

directories=("conduit")

for dir in "${directories[@]}"; do
    echo "initializing services in $dir"
    cd $dir || { echo "Directory $dir not found"; exit 1; }
    go run src/server.go || { echo "Failed to run server in $dir"; exit 1; }
    cd ..
done

echo "services initialized"
