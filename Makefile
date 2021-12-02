year = $(shell date +%Y)
day = $(shell date +'%-d')

new:
	mkdir -p $(year)/{input,solutions}; \
	./getinput.sh
	cp template.go $(year)/solutions/$(day).go; \