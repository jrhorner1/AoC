year = $(shell date +%Y)
day = $(shell date +'%-d')
repo = github.com/jrhorner1/AoC

new:
	mkdir -p $(year)/{input,go}; \
	./getinput.sh
	mkdir -p $(year)/go/day$(day); \
	cp template $(year)/go/day$(day)/$(day).go
	sed -i '' -e "s|package day0|package day$(day)|" $(year)/go/day$(day)/$(day).go; \
	sed -i '' \
		-e "s|^)|\t\"$(repo)/$(year)/go/day$(day)\"\n)|" \
		-e "s|^\tdefault:|\tcase $(day):\n\t\tinput, _ := ioutil.ReadFile(\"$(year)/input/$(day)\")\n\t\tfmt.Printf(\"\\\t%d Day %d solutions\\\nPart 1: %d\\\nPart 2\: %d\\\n\", \*year, \*day, day$(day).Puzzle(\&input, false), day$(day).Puzzle(\&input, true))\n\tdefault:|" \
		$(year)/go/$(year).go;