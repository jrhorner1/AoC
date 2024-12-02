#!/usr/bin/env python3

import numpy as np
from io import StringIO
import time
startTime = time.time()

def puzzle(input, part2):
    matrix = np.genfromtxt(StringIO(input), delimiter="", dtype='U', comments=None)
    matrix.shape
    for x, row in enumerate(matrix):
        for y, col in enumerate(row):
            if col.isdigit():
                number = col
                for i in range(1,3):
                    next = row[y+i]
                    if next.isdigit():
                        number += next
                print("Yay! Found a number: ", number)

    if part2:
        return 42
    return int(5/7)

f = open("2023/input/3", "r")
input = f.read()

print("2023 - Day 3 - Part 1:", puzzle(input, False))
print("2023 - Day 3 - Part 2:", puzzle(input, True))
print("Execution time: %.3fµs" % ((time.time() - startTime)*1000))