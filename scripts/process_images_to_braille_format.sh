#!/bin/bash
input=$1
output=$2
while IFS= read -r filepath
do
    fname="${filepath#*/*/*/}"
    plate="${fname%%.png}"
    eval echo "${plate}" ">>" "$output"
    cmd=(/usr/local/go/bin/dotmatrix -i -c 10 "<" "$filepath" ">>" "$output")
    eval "${cmd[@]}"
done < "$input"