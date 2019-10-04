package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInArray(t *testing.T) {
	t.Run("InArray", func(t *testing.T) {
		exists, index := InArray("s", []string{"a", "s", "D"})
		assert.Equal(t, true, exists)
		assert.Equal(t, 1, index)
	})
}

func TestInArrayInt(t *testing.T) {
	t.Run("InArrayInt", func(t *testing.T) {
		exists := InArrayInt(2, []int{4, 6, 2})
		assert.Equal(t, true, exists)
	})
}

func TestInArrayString(t *testing.T) {
	t.Run("InArrayString", func(t *testing.T) {
		exists := InArrayString("s", []string{"a", "s", "D"})
		assert.Equal(t, true, exists)
	})
}

func TestUniqueIntSlice(t *testing.T) {
	t.Run("UniqueIntSlice", func(t *testing.T) {
		newArray := UniqueIntSlice([]int{1, 2, 2, 5, 7, 8, 5})
		assert.Equal(t, newArray, []int{1, 2, 5, 7, 8})
	})
}

func TestRemoveItemInt(t *testing.T) {
	t.Run("RemoveItemInt", func(t *testing.T) {
		newArray := RemoveItemInt(6, []int{4, 6, 2})
		assert.Equal(t, newArray, []int{4, 2})
	})
}

func TestRemoveItemString(t *testing.T) {
	t.Run("RemoveItemString", func(t *testing.T) {
		newArray := RemoveItemString("bar", []string{"Foo", "bar", "meal"})
		assert.Equal(t, newArray, []string{"Foo", "meal"})
	})
}

func TestIsLetter(t *testing.T) {
	t.Run("IsLetter", func(t *testing.T) {
		isLetter := IsLetter("A")
		assert.Equal(t, true, isLetter)
	})

	t.Run("IsLetter", func(t *testing.T) {
		isLetter := IsLetter("0")
		assert.Equal(t, false, isLetter)
	})
}

func TestIsLetterAndDigit(t *testing.T) {
	t.Run("IsLetterAndDigit", func(t *testing.T) {
		isLetter := IsLetterAndDigit("A")
		assert.Equal(t, true, isLetter)
	})

	t.Run("IsLetterAndDigit", func(t *testing.T) {
		isLetter := IsLetterAndDigit("0")
		assert.Equal(t, true, isLetter)
	})

	t.Run("IsLetterAndDigit", func(t *testing.T) {
		isLetter := IsLetterAndDigit("%")
		assert.Equal(t, false, isLetter)
	})
}

func TestIsEnglishName(t *testing.T) {
	t.Run("IsEnglishName", func(t *testing.T) {
		isLetter := IsEnglishName("John")
		assert.Equal(t, true, isLetter)
	})

	t.Run("IsEnglishName", func(t *testing.T) {
		isLetter := IsEnglishName("Василий")
		assert.Equal(t, false, isLetter)
	})
}
