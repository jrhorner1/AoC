#!/usr/bin/env python3

import re
import time
startTime = time.time()

def puzzle(input, p2):
    numbers = [] 
    for line in input.splitlines():
        if p2:
            line = parseDigits(line)
        digits = re.findall(r'\d', line)
        numbers.append(int(digits[0] + digits[len(digits)-1]))
    return sum(numbers)

def parseDigits(line):
    digits = [ r"one", r"two", r"three", r"four", r"five", r"six", r"seven", r"eight", r"nine" ]
    calibrationStr = ""
    for i, c in enumerate(line):
        if c.isdigit():
            calibrationStr += c
            continue
        for j, digit in enumerate(digits):
            match = re.match(digit, line[i:])
            if match: 
                calibrationStr += str(j+1)
    return calibrationStr

f = open("2023/input/1", "r")
input = f.read()

print("2023 - Day 1 - Part 1:", puzzle(input, False))
print("2023 - Day 1 - Part 2:", puzzle(input, True))
print("Execution time: %.3fms" % ((time.time() - startTime)*1000))