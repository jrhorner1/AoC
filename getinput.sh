#!/bin/bash

# Usage function
_usage(){
        printf "%s\n" "$0 -s path/to/session/key -y [0-9]{4} -d [0-9]{2} ]"
        printf "\t%s\n" "-h|--help|-?  Print this help message."
        printf "\t%s\n" "-s|--session  Optional session file location. Default: [./session]"
        printf "\t%s\n" "-y|--year  4 digit year to retrieve input for. Required if not December."
        printf "\t%s\n" "-d|--day  2 digit day to retrieve input for. Required if not December."
        printf "\t%s\n" "-o|--output  Optional output filename. Default: [input]"
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
if [[ ${#sessionValue} != 96 ]]; then
    printf "Session hex malformed.\n"
    exit 2
fi

# Set some variables 
year=${year:=`date +%Y`}
day=${day:=`date +%d`}
outputPath=${year}/${day}
outputFile=${outputFile:="input"}
output=${outputPath}/${outputFile}

# Function to get the input file
_getInput(){
    local _day=`printf "${day}" | sed -e 's/^0//'`
    local url=https://adventofcode.com/${year}/day/${_day}/input 
    curl --cookie session=${sessionValue} -o ${output} ${url}
}

# Test if the overwrite flag is set
overwrite=${overwrite:=0}
if [[ ${overwrite} -eq 0 ]]; then
    # Validate the output path, create if necessary
    stat ${outputPath} > /dev/null 2>&1
    if [[ $? != 0 ]]; then
        mkdir -p ${outputPath}
        _getInput
    else 
        # Validate the output file status
        stat ${outputFile} > /dev/null 2>&1
        if [[ $? != 0 ]]; then
            _getInput
        else
            printf "File exists.\n"
            exit 3
        fi
    fi
else # Just send it (overwrite flag enabled)
    mkdir -p ${outputPath}
    _getInput
fi

