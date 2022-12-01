#!/bin/bash

# Usage function
_usage(){
        printf "%s\n" "$0 -s path/to/session/key -y [0-9]{4} -d [0-9]{2} ]"
        printf "\t%s\n" "-h|--help|-?  Print this help message."
        printf "\t%s\n" "-s|--session  Session key file location. Defaults to ./session"
        printf "\t%s\n" "-u|--useragent  User agent header to include with http request."
        printf "\t%s\n" "-y|--year  4 digit year to retrieve input for. Defaults to current year"
        printf "\t%s\n" "-d|--day  1 or 2 digit day to retrieve input for. Defaults to current day."
        printf "\t%s\n" "-o|--output  Output filename. Defaults to the configured day."
        printf "\t%s\n" "-f|--force  Force overwrite of existing input file."
		exit 0
}
# Print usage if no options are specified and it is not December
if [[ $# -eq 0 && `date +%m` != 12 ]]; then
    _usage
fi

# Parse cli options
while :; do
    case $1 in
        -h|--help)
            _usage
            ;;
        -s|--session)
            sessionFile=$2
            shift
            ;;
        -u|--useragent)
            userAgent=$2
        -y|--year)
            year=$2
            shift
            ;;
        -d|--day)
            day=$2
            shift
            ;;
        -o|--output)
            outputFile=$2
            shift
            ;;
        -f|--force)
            overwrite=1
            ;;
        *)
            break
            ;;
    esac
    shift
done

# Validate session file exists
sessionFile=${sessionFile:=./session}
stat ${sessionFile} > /dev/null 2>&1
if [[ $? != 0 ]]; then
    printf "Session file missing.\n"
    exit 1
fi
# Validate session hex length
sessionValue=`cat ${sessionFile}`
if [[ ${#sessionValue} != 128 ]]; then
    printf "Session hex malformed.\n"
    exit 2
fi

# Set some variables 
year=${year:=`date +%Y`}
day=${day:=`date +%-d`}
outputPath=${year}/input
outputFile=${outputFile:=${day}}
output=${outputPath}/${outputFile}
userAgent=${userAgent:="github.com/jrhorner1/AoC/getinput.sh by jrhorner@pm.me"}

# Function to get the input file
_getInput(){
    mkdir -p ${outputPath}
    curl --cookie session=${sessionValue} -A "${userAgent}" -o ${output} https://adventofcode.com/${year}/day/${day}/input 
}

# Validate the output, create if necessary, fetch input
overwrite=${overwrite:=0}
stat ${output} > /dev/null 2>&1
if [[ $? != 0 ]]; then
    _getInput
elif [[ ${overwrite} == 1 ]]; then
    _getInput
else
    printf "File exists.\n"
    exit 3
fi

