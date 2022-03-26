#!/bin/bash

git ls-files --others --ignored --exclude-standard | awk '{print}' | 
while read -r line
do
    if [[ $line == *"/ent/"* ]]; then
        echo "Delete: $line"
        rm $line
    fi
done


