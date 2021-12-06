## Advent of Code  

This is a repo for my [Advent of Code](https://adventofcode.com) attempts.

### Solutions
Puzzle solutions are meant to be run from the root of the repo. Without any arguments it will run the solution for the current day and year. 
```bash
go run main.go -y 2015 -d 2
```

### Utilities
Setup todays puzzle:
```bash
make new
``` 
Setup an older puzzle:
```bash
make new year=20XX day=X
``` 
Both of these commands will setup the directory structure, pull the input file, and create the barebones program from a template.  

Run the `getinput.sh` script without any options to pull the current days input. Use `-h` or `--help` to view the available options. This script requires your session key to work, so login to advent of code in a browser, open up web developer tools and copy the session cookie key (96 character hexidecimal) to a file named `session` at the root of this repo. 

### Updating old code
Previous years (specifically 19/20) have not been updated to work with the latest version of Go and as such some of the utility code does not work. Might get to fixing that at some point but for now, 2020 code has had the input paths updated where it was previously specified and 2019 code has not been modified at all so YMMV. 

### Archives
I archived my older repos and moved the code to this one, however they are still available as archives.
* [2020](https://github.com/jrhorner1/aoc2020)
* [2019](https://github.com/jrhorner1/aoc2019)




