#!/usr/bin/env python3

import time
startTime = time.time()

def puzzle(input, part2):
    if part2:
        return 42
    return int(5/7)

f = open("2023/input/1", "r")
input = f.read()

print("2023 - Day 1 - Part 1:", puzzle(input, False))
print("2023 - Day 1 - Part 2:", puzzle(input, True))
print("Execution time: %.3fµs" % ((time.time() - startTime)*1000))