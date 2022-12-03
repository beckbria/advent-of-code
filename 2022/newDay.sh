#!/bin/bash

if [ "$#" -ne 1 ]; then
    echo "Syntax: $0 <day>"
else
    cp -r 01/ $1
    sed -i "s/01/$1/g" $1/BUILD.bazel
    sed -i "s/2022\/01/2022\/$1/g" $1/main.go
fi