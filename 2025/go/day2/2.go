package day2

import (
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"
)

const DEBUG = false

func Puzzle(input *[]byte, part2 bool) int {
	if DEBUG {
		log.SetLevel(log.DebugLevel)
	}

	total := 0

	for _, id_range := range strings.Split(strings.TrimSpace(string(*input)), ",") {
		log.Debug(id_range)
		ids := strings.Split(strings.TrimSpace(id_range), "-")
		start, err := strconv.Atoi(ids[0])
		if err != nil {
			log.Error(err)
		}
		end, err := strconv.Atoi(ids[1])
		if err != nil {
			log.Error(err)
		}
		//log.Debugf("%d %d: %d", start, end, end-start)
	loop:
		for i := start; i <= end; i++ {
			id := strconv.Itoa(i)
			log.Debugf("id: %s", id)
			if part2 {
				// 10 digits max
				max := 10 / 2
				for n := 1; n <= max; n++ {
					if len(id)%n != 0 {
						continue
					}
					seq := id[0:n]
					log.Debugf("seq: %s", seq)
					for j := n; j <= len(id)-n; j += n {
						next := id[j : j+n]
						log.Debugf("next: %s", next)
						if string(next) != seq {
							break
						}
						log.Debugf("j: %d, len(id)-n: %d", j, len(id)-n)
						if j == len(id)-n {
							log.Debugf("invalid: %s", id)
							total += i
							continue loop
						}
					}
				}
			} else {
				// assume all ids with an odd number of digits are valid
				if len(id)%2 == 0 {
					front := id[0 : len(id)/2]
					back := id[len(id)/2:]
					//log.Debugf("%s: %s|%s", id, front, back)
					if front == back {
						log.Debugf("invalid: %s", id)
						total += i
					}
				}
			}
		}
	}
	return total
}
