package test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	sl "github.com/panter-dsd/stringlist/pkg/stringlist"
)

func TestStringList_SortEmpty(t *testing.T) {
	s := sl.StringList{}
	s.Sort()
	assert.Empty(t, s)
}

func TestStringList_SortSingleElement(t *testing.T) {
	a := sl.StringList{"a"}
	b := a
	b.Sort()
	assert.Equal(t, a, b)
}

func TestStringList_SortNotSorted(t *testing.T) {
	a := sl.StringList{"b", "a"}
	a.Sort()
	assert.Equal(t, sl.StringList{"a", "b"}, a)
}

func TestStringList_SortSorted(t *testing.T) {
	a := sl.StringList{"a", "b"}
	a.Sort()
	assert.Equal(t, sl.StringList{"a", "b"}, a)
}

func TestStringList_IsEmpty(t *testing.T) {
	assert.True(t, (&sl.StringList{}).IsEmpty())
	assert.False(t, (&sl.StringList{"a"}).IsEmpty())
}

func TestIsEqual(t *testing.T) {
	assert.True(t, sl.IsEqual(sl.StringList{}, sl.StringList{}))
	assert.True(t, sl.IsEqual(sl.StringList{"a"}, sl.StringList{"a"}))
	assert.True(t, sl.IsEqual(sl.StringList{"a", "b"}, sl.StringList{"a", "b"}))

	assert.False(t, sl.IsEqual(sl.StringList{"a"}, sl.StringList{"b"}))
	assert.False(t, sl.IsEqual(sl.StringList{"a", "b"}, sl.StringList{"a"}))
	assert.False(t, sl.IsEqual(sl.StringList{"a", "b"}, sl.StringList{"b", "a"}))
}

func TestStringList_Contains(t *testing.T) {
	assert.False(t, (&sl.StringList{}).Contains(""))
	assert.False(t, (&sl.StringList{}).Contains("a"))
	assert.False(t, (&sl.StringList{"b", "c"}).Contains("a"))

	assert.True(t, (&sl.StringList{"b", "c"}).Contains("c"))
	assert.True(t, (&sl.StringList{"c"}).Contains("c"))
}

func TestStringList_ForEach(t *testing.T) {
	a := sl.StringList{"a", "b", "c"}
	i := 0
	a.ForEach(func(s string) {
		assert.Equal(t, a[i], s)
		i++
	})
}
