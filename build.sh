#!/bin/sh

set -e

MY_DIR=$(dirname "$0")

cd "${MY_DIR}"
mkdir -p "artifacts"

echo "Linux"
GOARCH=amd64 GOOS=linux go build -o "artifacts/jenigma"

echo "Windows"
GOARCH=amd64 GOOS=windows go build -o "artifacts/jenigma.exe"

echo "Darwin"
GOARCH=amd64 GOOS=darwin go build -o "artifacts/jenigma_darwin"

echo "Build done"
