package day21

import (
	"strconv"
	"strings"
)

func Puzzle(input *[]byte, part2 bool) int {
	p1, p2 := parseInput(input)
	if part2 {
		p1.wins, p2.wins = quantumDirac(p1, p2)
		if p1.wins > p2.wins {
			return p1.wins
		} else {
			return p2.wins
		}
	}
	rolls := 0
	deterministicDirac(&p1, &p2, 100, &rolls)
	if p1.score > p2.score {
		return p2.score * rolls
	} else {
		return p1.score * rolls
	}
}

type player struct {
	position, score, wins int
}

func parseInput(input *[]byte) (player, player) {
	p1, p2 := player{0, 0, 0}, player{0, 0, 0}
	for i, line := range strings.Split(strings.TrimSpace(string(*input)), "\n") {
		startingPositions := strings.Split(line, ": ")
		if i == 0 {
			p1.position, _ = strconv.Atoi(startingPositions[1])
		} else {
			p2.position, _ = strconv.Atoi(startingPositions[1])
		}
	}
	return p1, p2
}

func deterministicDirac(p1, p2 *player, die int, rolls *int) {
	winningScore := 1000
	p := p1
	for p1.score < winningScore && p2.score < winningScore {
		sum := 0
		for rollCount := *rolls + 3; *rolls < rollCount; *rolls++ {
			die = wrap(die+1, 100)
			sum += die
		}
		p.position = wrap(p.position+sum, 10)
		p.score += p.position
		if p == p1 {
			p = p2
		} else {
			p = p1
		}
	}
}

func wrap(n, max int) int {
	if n%max == 0 {
		return max
	}
	return n % max
}

func quantumDirac(p1, p2 player) (int, int) {
	winningScore := 21
	if p1.score >= winningScore {
		return 1, 0
	} else if p2.score >= winningScore {
		return 0, 1
	}
	wins1, wins2 := 0, 0
	sumCombos := map[int]int{3: 1, 4: 3, 5: 6, 6: 7, 7: 6, 8: 3, 9: 1}
	for sum, combos := range sumCombos {
		np1 := p1
		np1.position = wrap(np1.position+sum, 10)
		np1.score += np1.position
		subWins2, subWins1 := quantumDirac(p2, np1)
		wins1, wins2 = wins1+combos*subWins1, wins2+combos*subWins2
	}
	return wins1, wins2
}
