year = $(shell date +%Y)
day = $(shell date +'%-d')

new:
	mkdir -p $(year)/{input,go}; \
	./getinput.sh
	mkdir -p $(year)/go/day$(day); \
	cp template.go $(year)/go/day$(day)/$(day).go
	sed -i '' -e "s/day0/day$(day)/" -e "s|0000/input/0|$(year)/input/$(day)|" $(year)/go/day$(day)/$(day).go; \