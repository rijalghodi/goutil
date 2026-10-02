package strutil

import "testing"

func TestReverse(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"empty string", "", ""},
		{"single char", "a", "a"},
		{"palindrome", "racecar", "racecar"},
		{"hello", "hello", "olleh"},
		{"with spaces", "hello world", "dlrow olleh"},
		{"unicode", "こんにちは", "はちにんこ"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Reverse(tt.input)
			if result != tt.expected {
				t.Errorf("Reverse(%q) = %q; want %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestTitle(t *testing.T) {
	if got := Title("hello world"); got != "Hello World" {
		t.Errorf("Title() = %q; want %q", got, "Hello World")
	}
}

func TestSlug(t *testing.T) {
	if got := Slug("Hello World!"); got != "hello-world" {
		t.Errorf("Slug() = %q; want %q", got, "hello-world")
	}
}

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  bool
	}{
		{"empty", "", true},
		{"single", "a", true},
		{"racecar", "racecar", true},
		{"hello", "hello", false},
		{"unicode", "あいういあ", true},
		{"unicode not", "あいう", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsPalindrome(tt.input); got != tt.want {
				t.Errorf("IsPalindrome(%q) = %v; want %v", tt.input, got, tt.want)
			}
		})
	}
}

func TestIsAnagram(t *testing.T) {
	tests := []struct {
		name string
		a, b string
		want bool
	}{
		{"empty", "", "", true},
		{"listen silent", "listen", "silent", true},
		{"unicode", "あい", "いあ", true},
		{"different length", "ab", "a", false},
		{"different chars", "abc", "abd", false},
		{"repeated", "aab", "aba", true},
		{"repeated mismatch", "aab", "abb", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsAnagram(tt.a, tt.b); got != tt.want {
				t.Errorf("IsAnagram(%q, %q) = %v; want %v", tt.a, tt.b, got, tt.want)
			}
		})
	}
}

func TestIsBlank(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{"empty", "", true},
		{"space", " ", true},
		{"tabs and newlines", "\t\n \r", true},
		{"not blank", " a ", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsBlank(tt.input)
			if result != tt.expected {
				t.Errorf("IsBlank(%q) = %v; want %v", tt.input, result, tt.expected)
			}
		})
	}
}
