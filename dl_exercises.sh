#!/bin/bash

while IFS= read -r line; do
  eval "$line"
done < "./dl.txt"