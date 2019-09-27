package utils

import (
	"reflect"
	"unicode"
)

// InArray это реализация in_array на go: https://codereview.stackexchange.com/questions/60074/in-array-in-go
func InArray(val interface{}, array interface{}) (exists bool, index int) {
	exists = false
	index = -1

	if reflect.TypeOf(array).Kind() == reflect.Slice {
		s := reflect.ValueOf(array)

		for i := 0; i < s.Len(); i++ {
			if reflect.DeepEqual(val, s.Index(i).Interface()) {
				index = i
				exists = true
				return
			}
		}
	}

	return
}

// InArrayInt returns true if val exists in array
func InArrayInt(val int, array []int) bool {

	for _, test := range array {
		if test == val {
			return true
		}
	}

	return false
}

// InArrayString returns true if val exists in array
func InArrayString(val string, array []string) bool {

	for _, test := range array {
		if test == val {
			return true
		}
	}

	return false
}

// UniqueIntSlice returns unique values in slice
func UniqueIntSlice(intSlice []int) []int {
	keys := make(map[int]bool)
	var list []int
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

// RemoveItemInt remove item from array
func RemoveItemInt(val int, array []int) (list []int) {

	for _, v := range array {
		if v == val {
			continue
		} else {
			list = append(list, v)
		}
	}

	return
}

// RemoveItemString remove item from array
func RemoveItemString(val string, array []string) (list []string) {

	for _, v := range array {
		if v == val {
			continue
		} else {
			list = append(list, v)
		}
	}

	return
}

// IsLetter vwhether s letter
func IsLetter(s string) bool {
	for _, r := range s {
		//if !unicode.IsLetter(r) {
		//	return false
		//}
		if (r < 'a' || r > 'z') && (r < 'A' || r > 'Z') {
			return false
		}
	}
	return true
}

// IsLatterAndDigit returns bool if s is Letter or digit
func IsLatterAndDigit(s string) bool {
	for _, r := range s {
		if (r < 'a' || r > 'z') && (r < 'A' || r > 'Z') && !unicode.IsDigit(r) {
			return false
		}
	}
	return true
}

// IsEnglishName returns true if all symbols is latin
func IsEnglishName(s string) bool {
	if s == "" {
		return false
	}
	for _, r := range s {
		if r == ' ' || r == '-' {
			continue
		}
		if (r < 'a' || r > 'z') && (r < 'A' || r > 'Z') {
			return false
		}
	}
	return true
}
