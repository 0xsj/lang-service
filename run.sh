#!/bin/sh

base_path="apps"  # Base path for service directories
directories=("auth" "conduit" "identity" "account")
pids=()  

for dir in "${directories[@]}"; do
    service_path="$base_path/$dir"
    echo "Initializing services in $service_path"
    (
        cd "$service_path" || { echo "Directory $service_path not found"; exit 1; }
        go run src/main.go || { echo "Failed to run server in $service_path"; exit 1; }
    ) &
    pids+=($!)  
done

for pid in "${pids[@]}"; do
    wait $pid || { echo "Server failed with PID $pid"; exit 1; }
done

echo "Services initialized"
