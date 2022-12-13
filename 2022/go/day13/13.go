package day13

import (
	"encoding/json"
	"fmt"
	"reflect"
	"sort"
	"strings"

	"github.com/sirupsen/logrus"
)

const dividerPackets string = `[[2]]
[[6]]`

func Puzzle(input *[]byte, part2 bool) int {
	//logrus.SetLevel(logrus.DebugLevel)
	var packets Packets
	for _, pair := range strings.Split(strings.TrimSpace(string(*input)), "\n\n") {
		for _, line := range strings.Split(strings.TrimSpace(pair), "\n") {
			var packet []any
			err := json.Unmarshal([]byte(line), &packet)
			if err != nil {
				logrus.Error(err)
			}
			packets = append(packets, &packet)
		}
	}
	if part2 {
		var dividers Packets
		for _, line := range strings.Split(strings.TrimSpace(dividerPackets), "\n") {
			var packet []any
			err := json.Unmarshal([]byte(line), &packet)
			if err != nil {
				logrus.Error(err)
			}
			packets = append(packets, &packet)
			dividers = append(dividers, &packet)
		}
		sort.Sort(packets)
		decoderKey := 1
		for _, divider := range dividers {
			for i, packet := range packets {
				if packet == divider {
					decoderKey *= i + 1
				}
			}
		}
		return decoderKey
	}
	sum := 0
	index := 1
	for i, j := 0, 1; j < len(packets); i, j = i+2, j+2 {
		if ok, err := inOrder(packets[i], packets[j]); err == nil && ok {
			sum += index
		}
		index++
	}
	return sum
}

type Packets []*[]any

func (p Packets) Len() int      { return len(p) }
func (p Packets) Swap(i, j int) { p[i], p[j] = p[j], p[i] }
func (p Packets) Less(i, j int) bool {
	if ok, err := inOrder(p[i], p[j]); err == nil && ok {
		return true
	}
	return false
}

func inOrder(left, right *[]any) (bool, error) {
	for i := 0; i < len(*left); i++ {
		if i >= len(*right) {
			return false, nil
		}
		leftType := reflect.TypeOf((*left)[i])
		rightType := reflect.TypeOf((*right)[i])
		if leftType == rightType {
			if leftType.Name() == "float64" {
				leftValue, rightValue := (*left)[i].(float64), (*right)[i].(float64)
				if leftValue < rightValue {
					return true, nil
				} else if leftValue > rightValue {
					return false, nil
				}
			} else {
				leftInner, rightInner := (*left)[i].([]any), (*right)[i].([]any)
				if result, err := inOrder(&leftInner, &rightInner); err == nil {
					return result, nil
				}
			}
		} else {
			if leftType.Name() == "float64" {
				leftInner, rightInner := []any{(*left)[i]}, (*right)[i].([]any)
				if result, err := inOrder(&leftInner, &rightInner); err == nil {
					return result, nil
				}
			} else if rightType.Name() == "float64" {
				leftInner, rightInner := (*left)[i].([]any), []any{(*right)[i]}
				if result, err := inOrder(&leftInner, &rightInner); err == nil {
					return result, nil
				}
			}
		}
	}
	if len(*left) < len(*right) {
		return true, nil
	}
	return false, fmt.Errorf("")
}
