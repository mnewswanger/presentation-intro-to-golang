#!/bin/bash
DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )/"

go run $DIR/multi-safe/multi-safe.go
