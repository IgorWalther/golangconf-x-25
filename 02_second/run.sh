#!/bin/bash

set -e

for i in {1..100000}; do
    echo "Run $i"
    go run main.go
done

echo "No bugs found in 100000 runs"
