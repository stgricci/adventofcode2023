#!/bin/bash

echo "Part 1 --"
cat puzzle_input.txt | jq --raw-input --null-input --from-file part1.jq
echo "--"
echo "Part 2 --"
cat puzzle_input.txt | jq --raw-input --null-input --from-file part2.jq
