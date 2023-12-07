package day7

import (
	"sort"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"
)

var (
	cardStrength = map[rune]int{'2': 1, '3': 2, '4': 3, '5': 4, '6': 5, '7': 6, '8': 7, '9': 8, 'T': 9, 'J': 10, 'Q': 11, 'K': 12, 'A': 13}
	typeStrength = map[string]int{"5K": 7, "4K": 6, "FH": 5, "3K": 4, "2P": 3, "1P": 2, "HC": 1}
)

type hand struct {
	cards, hType string
	bid          int
}

type byStrength []hand

func (h byStrength) Len() int { return len(h) }
func (h byStrength) Less(i, j int) bool {
	if h[i].hType == h[j].hType {
		for k := 0; k < 5; k++ {
			a, b := rune(h[i].cards[k]), rune(h[j].cards[k])
			if a != b {
				return cardStrength[a] < cardStrength[b]
			}
		}
	}
	return typeStrength[h[i].hType] < typeStrength[h[j].hType]
}
func (h byStrength) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func Puzzle(input *[]byte, part2 bool) int {
	//log.SetLevel(log.DebugLevel)
	hands := []hand{}
	for _, line := range strings.Split(strings.TrimSpace(string(*input)), "\n") {
		content := strings.Fields(line)
		cards, hType := content[0], ""
		bid, _ := strconv.Atoi(content[1])
		cardQuantity := map[rune]int{}
		for _, card := range cards {
			cardQuantity[card] += 1
		}
		if part2 {
			cardStrength['J'] = 0
			if wild, ok := cardQuantity['J']; ok && len(cardQuantity) > 1 {
				for k := range cardQuantity {
					if k != 'J' {
						cardQuantity[k] += wild
					}
				}
				delete(cardQuantity, 'J')
			}
		}
		switch len(cardQuantity) {
		case 5: // high card
			hType = "HC"
		case 4: // one pair
			hType = "1P"
		case 3: // three of a kind or two pair
			hType = "2P"
			for _, quantity := range cardQuantity {
				if quantity == 3 {
					hType = "3K"
				}
			}
		case 2: // full house or 4 of a kind
			hType = "FH"
			for _, quantity := range cardQuantity {
				if quantity == 4 {
					hType = "4K"
				}
			}
		case 1: // five of a kind
			hType = "5K"
		}
		newHand := hand{cards, hType, bid}
		hands = append(hands, newHand)
	}
	// sort hands by strength ascending
	sort.Sort(byStrength(hands))
	log.Debug("Sorted Hands: ", hands)
	winnings := 0
	for rank, h := range hands {
		winnings += (h.bid * (rank + 1))
	}
	return winnings
}
