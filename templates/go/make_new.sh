#!/bin/bash

year=${year:=`date +%Y`}
yr=`echo ${year} | sed -E 's/20([0-9]{2})/\1/'`
day=${day:=`date +%-d`}
repo="github.com/jrhorner1/AoC"
sed_opts="-i"
if [[ ! ${OSTYPE} == "linux-gnu" ]]; then
    sed_opts="-i ''"
fi

mkdir -p ${year}/{input,go}
mkdir -p ${year}/go/day${day}

if [[ ! -f ${year}/go/${year}.go ]]; then 
    cp templates/go/year ${year}/go/${year}.go 
    sed ${sed_opts} -e "s|package y00|package y${yr}|" ${year}/go/${year}.go
    sed ${sed_opts} \
        -e "s|^)|\ty${yr} \"${repo}/${year}/go\"\n)|" \
        -e "s|^\tdefault:|\tcase ${year}:\n\t\ty${yr}.Run\(\&year, \&day\)\n\tdefault:|" \
        main.go
fi
if [[ ! -f ${year}/go/day${day}/${day}.go ]]; then
    cp templates/go/day ${year}/go/day${day}/${day}.go
    sed ${sed_opts} -e "s|package day0|package day${day}|" ${year}/go/day${day}/${day}.go
    sed ${sed_opts} \
        -e "s|^)|\t\"${repo}/${year}/go/day${day}\"\n)|" \
        -e "s|^\tdefault:|\tcase ${day}:\n\t\tinput, _ := ioutil.ReadFile(\"${year}/input/${day}\")\n\t\tfmt.Printf(\"\\\t%d Day %d solutions\\\nPart 1: %d\\\nPart 2\: %d\\\n\", \*year, \*day, day${day}.Puzzle(\&input, false), day${day}.Puzzle(\&input, true))\n\tdefault:|" \
        ${year}/go/${year}.go
fi
if [[ ! -f ${year}/go/day${day}/${day}_test.go ]]; then
    cp templates/go/day_test ${year}/go/day${day}/${day}_test.go
    sed ${sed_opts} -e "s|package day0|package day${day}|" ${year}/go/day${day}/${day}_test.go
fi
