package day13

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"

	"github.com/sirupsen/logrus"
)

func Puzzle(input *[]byte, part2 bool) int {
	//logrus.SetLevel(logrus.DebugLevel)
	sum := 0
	for i, pair := range strings.Split(strings.TrimSpace(string(*input)), "\n\n") {
		var left, right []any
		for j, line := range strings.Split(strings.TrimSpace(pair), "\n") {
			var packet *[]any
			if j == 0 {
				packet = &left
			} else {
				packet = &right
			}
			err := json.Unmarshal([]byte(line), packet)
			if err != nil {
				logrus.Error(err)
			}
		}
		if ok, err := inOrder(&left, &right); err == nil && ok {
			sum += i + 1
		}
	}
	if part2 {
		return 42
	}
	return sum
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
