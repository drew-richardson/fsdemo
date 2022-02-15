#!/bin/sh
set -e

SCRIPT_DIR="$(cd "$(dirname "$0")"; pwd)"
cd $SCRIPT_DIR

go run cmd/embed/main.go

go run cmd/dirfs/main.go

tar -cf static.tar -C internal/data static
go run cmd/tarfs/main.go
