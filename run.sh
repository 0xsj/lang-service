#!/bin/sh

directories=("auth" "conduit")
pids=()  

for dir in "${directories[@]}"; do
    echo "Initializing services in $dir"
    (
        cd "$dir" || { echo "Directory $dir not found"; exit 1; }
        go run src/main.go || { echo "Failed to run server in $dir"; exit 1; }
    ) &
    pids+=($!)  
done

for pid in "${pids[@]}"; do
    wait $pid || { echo "Server failed with PID $pid"; exit 1; }
done

echo "Services initialized"
