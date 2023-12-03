#!/bin/bash

function usage
{
    echo "Creates a directory for a new day of Advent of Code"
    echo "Usage: ./newDay.sh dayNumber" 
    exit 1
}

# The single parameter must be a number. 
if [[ "$#" -ne 1 ]] || ! [[ "$1" =~ ^[0-9]+$ ]]; then
    usage
fi

if [ -d "$1" ]; then
    echo "Directory $1 already exists"
    exit 2
fi

cp -r template $1
for f in BUILD.bazel main.kt; do
   sed -i "s/{{DAYPLACEHOLDER}}/$1/g" $1/$f;
done
