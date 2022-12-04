#!/bin/bash

if [ "$#" -ne 1 ]; then
    echo "Syntax: $0 <day>"
else
    cp -r 03/ $1
    sed -i "s/03/$1/g" $1/BUILD.bazel
    sed -i "s/2022\/03/2022\/$1/g" $1/main.go
fi