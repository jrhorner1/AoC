## Advent of Code  

This is a repo for my [Advent of Code](https://adventofcode.com) attempts.

### Running the code

2021 code is meant to be run from the root of the repo. 

```bash
go run 2021/solutions/1.go
```

### Retrieving input

Run the `getinput.sh` script without any options to pull the current days input. Use `-h` or `--help` to view the available options. 

### Cloning this repo

After cloning this repo, the submodule folders will be empty. See https://git-scm.com/book/en/v2/Git-Tools-Submodules for more info.  

Run the following commands to get the missing files:  

```bash
git submodule init
git submodule update
```