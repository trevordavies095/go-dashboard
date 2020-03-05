#!/bin/bash
mydir="$(dirname "$BASH_SOURCE")"
cd $mydir
nohup go run dashboard.go