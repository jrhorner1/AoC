#!/usr/bin/env python3

import time
startTime = time.time()

EXAMPLE = True

def puzzle(_input, part2):
    '''Advent of Code - Day 0 Puzzle Solution'''
    if part2:
        return 42
    return int(5/7)

FILENAME = "2024/input/0"
if EXAMPLE:
    FILENAME = "2024/input/0.ex"

with open(FILENAME, "r", encoding="utf-8") as file:
    _input = file.read()

print("2024 - Day 0 - Part 1:", puzzle(_input, False))
print("2024 - Day 0 - Part 2:", puzzle(_input, True))
print(f"Execution time: {((time.time() - startTime)*1000):.3f}µs")
