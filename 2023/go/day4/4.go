package day4

import (
	"strings"

	log "github.com/sirupsen/logrus"
)

type card struct {
	id, matches, points, quantity int
}

func Puzzle(input *[]byte, part2 bool) int {
	//log.SetLevel(log.DebugLevel)
	cards := []card{}
	for i, line := range strings.Split(strings.TrimSpace(string(*input)), "\n") {
		if len(line) == 0 {
			continue
		}
		newCard := card{id: i, matches: 0, points: 0, quantity: 1}
		cardSlice := strings.Split(strings.ReplaceAll(line, "  ", " "), " | ")
		log.Debug(cardSlice)
		win := strings.Split(cardSlice[0], " ")
		log.Debug(win)
		have := strings.Split(cardSlice[1], " ")
		log.Debug(have)
		for i, n := range win {
			if i == 0 || i == 1 {
				continue
			}
			for _, h := range have {
				if h == n {
					log.Debug("Match! ", h, " ", n)
					newCard.matches++
					if newCard.points == 0 {
						newCard.points = 1
					} else {
						newCard.points *= 2
					}
				}
			}
		}
		log.Debug("Card ", win[1], " total points ", newCard.points)
		cards = append(cards, newCard)
	}
	sum, totalCards := 0, 0
	for i, c := range cards {
		sum += c.points
		log.Debug("Card ", c.id, " won ", c.matches, " additional cards.")
		for add := 0; add < c.quantity; add++ {
			for j := 1; j <= c.matches && i+j < len(cards); j++ {
				cards[i+j].quantity++
				log.Debug("Now have ", cards[i+j].quantity, " of card ", cards[i+j].id)
			}
		}
	}
	if part2 {
		for _, c := range cards {
			totalCards += c.quantity
		}
		return totalCards
	}
	return sum
}
