package str

import (
	"github.com/gosimple/slug"
)

// Reverse returns a new string with the characters in reverse order.
func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// IsBlank returns true if the string is empty or contains only whitespace.
func IsBlank(s string) bool {
	for _, r := range s {
		if r != ' ' && r != '\t' && r != '\n' && r != '\r' {
			return false
		}
	}
	return true
}

// Slug returns a URL-safe slug for s.
func Slug(s string) string {
	return slug.Make(s)
}

// IsPalindrome returns true if s reads the same forwards and backwards.
func IsPalindrome(s string) bool {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		if runes[i] != runes[j] {
			return false
		}
	}
	return true
}

// IsAnagram returns true if a and b contain the same characters with the same frequencies.
func IsAnagram(a, b string) bool {
	ra, rb := []rune(a), []rune(b)
	if len(ra) != len(rb) {
		return false
	}
	counts := make(map[rune]int, len(ra))
	for _, r := range ra {
		counts[r]++
	}
	for _, r := range rb {
		counts[r]--
		if counts[r] < 0 {
			return false
		}
	}
	return true
}
