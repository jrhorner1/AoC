package day2

import (
	"fmt"
	"strings"
)

var (
	winMap = map[string]string{
		"rock":     "scissors",
		"paper":    "rock",
		"scissors": "paper",
	}
	formMap = map[string]string{
		"A": "rock",
		"B": "paper",
		"C": "scissors",
		"X": "rock",
		"Y": "paper",
		"Z": "scissors",
	}
	actionMap = map[string]string{
		"X": "lose",
		"Y": "draw",
		"Z": "win",
	}
)

type RPS struct {
	form  string
	score int
}

func Puzzle(input *[]byte, part2 bool) int {
	player := RPS{"", 0}
	for _, round := range strings.Split(strings.TrimSpace(string(*input)), "\n") {
		play := strings.Split(strings.TrimSpace(round), " ")
		opponent := RPS{formMap[play[0]], 0}
		if part2 {
			player.action(play[1], &opponent)
		} else {
			player.form = formMap[play[1]]
		}

		opponent.addFormPoints()
		player.addFormPoints()

		addWinPoints(&opponent, &player)
	}
	return player.score
}

func (rps *RPS) action(play string, opponent *RPS) {
	switch actionMap[play] {
	case "lose":
		rps.form = winMap[opponent.form]
	case "draw":
		rps.form = opponent.form
	case "win":
		for k, v := range winMap {
			if v == opponent.form {
				rps.form = k
			}
		}
	}
}

func (rps *RPS) addFormPoints() {
	formPoints := map[string]int{
		"rock":     1,
		"paper":    2,
		"scissors": 3,
	}
	rps.score += formPoints[rps.form]
}

func addWinPoints(opponent, player *RPS) {
	if winMap[player.form] == opponent.form {
		// player won
		player.score += 6
	} else if winMap[opponent.form] == player.form {
		// opponent won
		opponent.score += 6
	} else if player.form == opponent.form {
		// draw
		player.score += 3
		opponent.score += 3
	} else {
		// something very wrong happened
		fmt.Println("How did you get here?")
	}
}
