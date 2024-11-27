#!/bin/bash

path="$*"

if [ -z "$path" ]; then
    echo "Please provide a folder name"
    exit
fi

path=$(echo "$path" | sed 's/ /-/g' | sed 's/\.//g')
mkdir "$path"

touch "$path/$path.go"

nvim "$path/$path.go"
