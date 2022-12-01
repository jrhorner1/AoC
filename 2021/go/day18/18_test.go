package day18

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
)

func Test_explode(t *testing.T) {
	input := []byte(`[[[[[9,8],1],2],3],4]
[7,[6,[5,[4,[3,2]]]]]
[[6,[5,[4,[3,2]]]],1]
[[3,[2,[1,[7,3]]]],[6,[5,[4,[3,2]]]]]
[[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]]
`)
	answers := []SnailfishSlice{
		{{0, 4}, {9, 4}, {2, 3}, {3, 2}, {4, 1}},                                 // [[[[0,9],2],3],4]
		{{7, 1}, {6, 2}, {5, 3}, {7, 4}, {0, 4}},                                 // [7,[6,[5,[7,0]]]]
		{{6, 2}, {5, 3}, {7, 4}, {0, 4}, {3, 1}},                                 // [[6,[5,[7,0]]],3]
		{{3, 2}, {2, 3}, {8, 4}, {0, 4}, {9, 2}, {5, 3}, {4, 4}, {3, 5}, {2, 5}}, // [[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]]
		{{3, 2}, {2, 3}, {8, 4}, {0, 4}, {9, 2}, {5, 3}, {7, 4}, {0, 4}},         // [[3,[2,[8,0]]],[9,[5,[7,0]]]]
	}
	numbers := parseInput(&input)
	for i := range numbers {
		numbers[i].Explode()
	}
	for i := range numbers {
		if !reflect.DeepEqual(numbers[i], answers[i]) {
			t.Errorf("got %d, wanted %d", numbers[i], answers[i])
		}
	}
}

func Test_split(t *testing.T) {
	input := []byte(`[[[[3,11],2],3],4]
`)
	answers := []SnailfishSlice{
		{{3, 4}, {5, 5}, {6, 5}, {2, 3}, {3, 2}, {4, 1}}, // [[[[3,[5,6]],2],3],4]
	}
	numbers := parseInput(&input)
	for i := range numbers {
		numbers[i].Split()
	}
	for i := range numbers {
		if !reflect.DeepEqual(numbers[i], answers[i]) {
			t.Errorf("got %d, wanted %d", numbers[i], answers[i])
		}
	}
}

func Test_reduce(t *testing.T) {
	input := []byte(`[[[[[4,3],4],4],[7,[[8,4],9]]],[1,1]]
`)
	answers := []SnailfishSlice{
		{{0, 4}, {7, 4}, {4, 3}, {7, 4}, {8, 4}, {6, 4}, {0, 4}, {8, 2}, {1, 2}}, // [[[[0,7],4],[[7,8],[6,0]]],[8,1]]
	}
	numbers := parseInput(&input)
	for i := range numbers {
		numbers[i].Reduce()
	}
	for i := range numbers {
		if !reflect.DeepEqual(numbers[i], answers[i]) {
			t.Errorf("got %d, wanted %d", numbers[i], answers[i])
		}
	}
}

func Test_add(t *testing.T) {
	inputs := [][]byte{
		[]byte(`[1,1]
[2,2]
[3,3]
[4,4]
`),
		[]byte(`
[1,1]
[2,2]
[3,3]
[4,4]
[5,5]
`),
		[]byte(`[1,1]
[2,2]
[3,3]
[4,4]
[5,5]
[6,6]
`),
		[]byte(`[[[0,[4,5]],[0,0]],[[[4,5],[2,6]],[9,5]]]
[7,[[[3,7],[4,3]],[[6,3],[8,8]]]]
[[2,[[0,8],[3,4]]],[[[6,7],1],[7,[1,6]]]]
[[[[2,4],7],[6,[0,5]]],[[[6,8],[2,8]],[[2,1],[4,5]]]]
[7,[5,[[3,8],[1,4]]]]
[[2,[2,2]],[8,[8,1]]]
[2,9]
[1,[[[9,3],9],[[9,0],[0,7]]]]
[[[5,[7,4]],7],1]
[[[[4,2],2],6],[8,7]]
`),
	}
	answers := []SnailfishSlice{
		{{1, 4}, {1, 4}, {2, 4}, {2, 4}, {3, 3}, {3, 3}, {4, 2}, {4, 2}},                                                 // [[[[1,1],[2,2]],[3,3]],[4,4]]
		{{3, 4}, {0, 4}, {5, 4}, {3, 4}, {4, 3}, {4, 3}, {5, 2}, {5, 2}},                                                 // [[[[3,0],[5,3]],[4,4]],[5,5]]
		{{5, 4}, {0, 4}, {7, 4}, {4, 4}, {5, 3}, {5, 3}, {6, 2}, {6, 2}},                                                 // [[[[5,0],[7,4]],[5,5]],[6,6]]
		{{8, 4}, {7, 4}, {7, 4}, {7, 4}, {8, 4}, {6, 4}, {7, 4}, {7, 4}, {0, 4}, {7, 4}, {6, 4}, {6, 4}, {8, 3}, {7, 3}}, // [[[[8,7],[7,7]],[[8,6],[7,7]]],[[[0,7],[6,6]],[8,7]]]
	}
	sums := []SnailfishSlice{}
	for _, input := range inputs {
		numbers, sum := parseInput(&input), SnailfishSlice{}
		for i, number := range numbers {
			if i == 0 {
				sum = number
				continue
			}
			sum.Add(&number)
		}
		sums = append(sums, sum)
	}
	for i := range sums {
		if !reflect.DeepEqual(sums[i], answers[i]) {
			t.Errorf("got %d, wanted %d", sums[i], answers[i])
		}
	}
}

func Test_magnitude(t *testing.T) {
	input := []byte(`[[1,2],[[3,4],5]]
[[[[0,7],4],[[7,8],[6,0]]],[8,1]]
[[[[1,1],[2,2]],[3,3]],[4,4]]
[[[[3,0],[5,3]],[4,4]],[5,5]]
[[[[5,0],[7,4]],[5,5]],[6,6]]
[[[[8,7],[7,7]],[[8,6],[7,7]]],[[[0,7],[6,6]],[8,7]]]
`)
	answers := []int{143, 1384, 445, 791, 1137, 3488}
	numbers := parseInput(&input)
	for i := range numbers {
		mag := numbers[i].Magnitude()
		if mag != answers[i] {
			t.Errorf("got %d, wanted %d", mag, answers[i])
		}
	}
}

func Test_largestSum(t *testing.T) {
	input := []byte(`[[[0,[5,8]],[[1,7],[9,6]]],[[4,[1,2]],[[1,4],2]]]
[[[5,[2,8]],4],[5,[[9,9],0]]]
[6,[[[6,2],[5,6]],[[7,6],[4,7]]]]
[[[6,[0,7]],[0,9]],[4,[9,[9,0]]]]
[[[7,[6,4]],[3,[1,3]]],[[[5,5],1],9]]
[[6,[[7,3],[3,2]]],[[[3,8],[5,7]],4]]
[[[[5,4],[7,7]],8],[[8,3],8]]
[[9,3],[[9,9],[6,[4,9]]]]
[[2,[[7,7],7]],[[5,8],[[9,3],[0,2]]]]
[[[[5,2],5],[8,[3,7]]],[[5,[7,5]],[4,4]]]
`)
	answer := 3993
	numbers, mags := parseInput(&input), []int{}
	// loopA:
	for _, a := range numbers {
	loopB:
		for _, b := range numbers {
			if reflect.DeepEqual(a, b) {
				continue loopB
			}
			a.Add(&b)
			mags = append(mags, a.Magnitude())
		}
	}
	sort.Ints(mags)
	fmt.Println(len(mags), mags)
	if mags[len(mags)-1] != answer {
		t.Errorf("got %d, wanted %d", mags[len(mags)-1], answer)
	}
}
