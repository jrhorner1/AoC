/*
Copyright The Kubernetes Authors.
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Modified from https://github.com/kubernetes/apimachinery/blob/v0.25.4/pkg/util/sets/string.go

package sets

import (
	"reflect"
	"sort"
)

// sets.Rune is a set of runes, implemented via map[rune]struct{} for minimal memory consumption.
type Rune map[rune]Empty

// NewRune creates a Rune from a list of values.
func NewRune(items ...rune) Rune {
	ss := make(Rune, len(items))
	ss.Insert(items...)
	return ss
}

// RuneKeySet creates a Rune from a keys of a map[rune](? extends interface{}).
// If the value passed in is not actually a map, this will panic.
func RuneKeySet(theMap interface{}) Rune {
	v := reflect.ValueOf(theMap)
	ret := Rune{}

	for _, keyValue := range v.MapKeys() {
		ret.Insert(keyValue.Interface().(rune))
	}
	return ret
}

// Insert adds items to the set.
func (s Rune) Insert(items ...rune) Rune {
	for _, item := range items {
		s[item] = Empty{}
	}
	return s
}

// Delete removes all items from the set.
func (s Rune) Delete(items ...rune) Rune {
	for _, item := range items {
		delete(s, item)
	}
	return s
}

// Has returns true if and only if item is contained in the set.
func (s Rune) Has(item rune) bool {
	_, contained := s[item]
	return contained
}

// HasAll returns true if and only if all items are contained in the set.
func (s Rune) HasAll(items ...rune) bool {
	for _, item := range items {
		if !s.Has(item) {
			return false
		}
	}
	return true
}

// HasAny returns true if any items are contained in the set.
func (s Rune) HasAny(items ...rune) bool {
	for _, item := range items {
		if s.Has(item) {
			return true
		}
	}
	return false
}

// Clone returns a new set which is a copy of the current set.
func (s Rune) Clone() Rune {
	result := make(Rune, len(s))
	for key := range s {
		result.Insert(key)
	}
	return result
}

// Difference returns a set of objects that are not in s2
// For example:
// s1 = {a1, a2, a3}
// s2 = {a1, a2, a4, a5}
// s1.Difference(s2) = {a3}
// s2.Difference(s1) = {a4, a5}
func (s Rune) Difference(s2 Rune) Rune {
	result := NewRune()
	for key := range s {
		if !s2.Has(key) {
			result.Insert(key)
		}
	}
	return result
}

// Union returns a new set which includes items in either s1 or s2.
// For example:
// s1 = {a1, a2}
// s2 = {a3, a4}
// s1.Union(s2) = {a1, a2, a3, a4}
// s2.Union(s1) = {a1, a2, a3, a4}
func (s1 Rune) Union(s2 Rune) Rune {
	result := s1.Clone()
	for key := range s2 {
		result.Insert(key)
	}
	return result
}

// Intersection returns a new set which includes the item in BOTH s1 and s2
// For example:
// s1 = {a1, a2}
// s2 = {a2, a3}
// s1.Intersection(s2) = {a2}
func (s1 Rune) Intersection(s2 Rune) Rune {
	var walk, other Rune
	result := NewRune()
	if s1.Len() < s2.Len() {
		walk = s1
		other = s2
	} else {
		walk = s2
		other = s1
	}
	for key := range walk {
		if other.Has(key) {
			result.Insert(key)
		}
	}
	return result
}

// IsSuperset returns true if and only if s1 is a superset of s2.
func (s1 Rune) IsSuperset(s2 Rune) bool {
	for item := range s2 {
		if !s1.Has(item) {
			return false
		}
	}
	return true
}

// Equal returns true if and only if s1 is equal (as a set) to s2.
// Two sets are equal if their membership is identical.
// (In practice, this means same elements, order doesn't matter)
func (s1 Rune) Equal(s2 Rune) bool {
	return len(s1) == len(s2) && s1.IsSuperset(s2)
}

type sortableSliceOfRune []rune

func (s sortableSliceOfRune) Len() int           { return len(s) }
func (s sortableSliceOfRune) Less(i, j int) bool { return lessRune(s[i], s[j]) }
func (s sortableSliceOfRune) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

// List returns the contents as a sorted rune slice.
func (s Rune) List() []rune {
	res := make(sortableSliceOfRune, 0, len(s))
	for key := range s {
		res = append(res, key)
	}
	sort.Sort(res)
	return []rune(res)
}

// UnsortedList returns the slice with contents in random order.
func (s Rune) UnsortedList() []rune {
	res := make([]rune, 0, len(s))
	for key := range s {
		res = append(res, key)
	}
	return res
}

// Returns a single element from the set.
func (s Rune) PopAny() (rune, bool) {
	for key := range s {
		s.Delete(key)
		return key, true
	}
	var zeroValue rune
	return zeroValue, false
}

// Len returns the size of the set.
func (s Rune) Len() int {
	return len(s)
}

func lessRune(lhs, rhs rune) bool {
	return lhs < rhs
}
