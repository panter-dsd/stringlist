package stringlist

import (
	"sort"
	"strings"
)

// StringList ...
type StringList []string

// Sort ...
func (s *StringList) Sort() *StringList {
	sort.Strings(*s)
	return s
}

// IsEmpty ...
func (s *StringList) IsEmpty() bool {
	return len(*s) == 0
}

func (s *StringList) Size() int {
	return len(*s)
}

// IsEqual ...
func IsEqual(a, b StringList) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

// Contains ...
func (s *StringList) Contains(str string) bool {
	for _, e := range *s {
		if e == str {
			return true
		}
	}
	return false
}

// ForEach run function from each string
func (s *StringList) ForEach(f func(s string)) {
	for _, str := range *s {
		f(str)
	}
}

func (s *StringList) At(index int) string {
	if index >= 0 && index < s.Size() {
		return (*s)[index]
	}
	return ""
}

func NewFromString(s, separator string) *StringList {
	sl := StringList(strings.Split(s, separator))
	return &sl
}
