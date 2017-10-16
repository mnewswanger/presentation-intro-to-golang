#!/bin/bash
DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )/"

cd $DIR/multi-race
go test -v --race
cd $DIR
