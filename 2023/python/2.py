#!/usr/bin/env python3

import time
startTime = time.time()

def puzzle(input, part2):
    totalCubes = { "red": 12, "green": 13, "blue": 14 }
    possibleSum, powerSum = 0, 0
    for line in input.splitlines():
        content = line.split(": ")
        id = int(content[0].split(" ")[1])
        minCubes = {"red": 0, "green": 0, "blue": 0}
        impossible, power = False, 1
        sets = content[1].split("; ")
        for set in sets:
            for cube in set.split(", "):
                cube = cube.split(" ")
                if not impossible and int(cube[0]) > totalCubes[cube[1]]:
                    impossible = True
                if minCubes[cube[1]] < int(cube[0]):
                    minCubes[cube[1]] = int(cube[0])
        if not impossible:
            possibleSum += id
        for color in minCubes:
            power *= minCubes[color]
        powerSum += power
    if part2:
        return powerSum
    return possibleSum

f = open("2023/input/2", "r")
input = f.read()

print("2023 - Day 2 - Part 1:", puzzle(input, False))
print("2023 - Day 2 - Part 2:", puzzle(input, True))
print("Execution time: %.3fµs" % ((time.time() - startTime)*1000))