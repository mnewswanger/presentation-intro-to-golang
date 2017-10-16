#!/bin/bash
DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )/"

cd $DIR/multi-safe
go test -v --race
cd $DIR
