package stringlist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringList_SortEmpty(t *testing.T) {
	s := StringList{}
	s.Sort()
	assert.Empty(t, s)
}

func TestStringList_SortSingleElement(t *testing.T) {
	a := StringList{"a"}
	b := a
	b.Sort()
	assert.Equal(t, a, b)
}

func TestStringList_SortNotSorted(t *testing.T) {
	a := StringList{"b", "a"}
	a.Sort()
	assert.Equal(t, StringList{"a", "b"}, a)
}

func TestStringList_SortSorted(t *testing.T) {
	a := StringList{"a", "b"}
	a.Sort()
	assert.Equal(t, StringList{"a", "b"}, a)
}

func TestStringList_IsEmpty(t *testing.T) {
	assert.True(t, (&StringList{}).IsEmpty())
	assert.False(t, (&StringList{"a"}).IsEmpty())
}

func TestIsEqual(t *testing.T) {
	assert.True(t, IsEqual(StringList{}, StringList{}))
	assert.True(t, IsEqual(StringList{"a"}, StringList{"a"}))
	assert.True(t, IsEqual(StringList{"a", "b"}, StringList{"a", "b"}))

	assert.False(t, IsEqual(StringList{"a"}, StringList{"b"}))
	assert.False(t, IsEqual(StringList{"a", "b"}, StringList{"a"}))
	assert.False(t, IsEqual(StringList{"a", "b"}, StringList{"b", "a"}))
}

func TestStringList_Contains(t *testing.T) {
	assert.False(t, (&StringList{}).Contains(""))
	assert.False(t, (&StringList{}).Contains("a"))
	assert.False(t, (&StringList{"b", "c"}).Contains("a"))

	assert.True(t, (&StringList{"b", "c"}).Contains("c"))
	assert.True(t, (&StringList{"c"}).Contains("c"))
}

func TestStringList_ForEach(t *testing.T) {
	a := StringList{"a", "b", "c"}
	i := 0
	a.ForEach(func(s string) {
		assert.Equal(t, a[i], s)
		i++
	})
}

func TestStringList_FromString(t *testing.T) {
	s := "some string with separator space"
	l := NewFromString(s, " ")
	assert.Equal(t, 5, l.Size())
	assert.Equal(t, StringList{"some", "string", "with", "separator", "space"}, *l)
}
