package strings

import "testing"

func TestKMP(t *testing.T) {
	tests := []struct {
		name      string
		text      string
		substring string
		expected  int
	}{
		{"found at start", "hello world", "hello", 0},
		{"found at end", "hello world", "world", 6},
		{"found in middle", "abcdefgh", "cde", 2},
		{"not found", "hello world", "xyz", -1},
		{"empty substring", "hello", "", 0},
		{"empty text", "", "hello", -1},
		{"both empty", "", "", 0},
		{"substring equals text", "abc", "abc", 0},
		{"substring longer than text", "ab", "abc", -1},
		{"repeated pattern", "aababab", "abab", 1},
		{"single char found", "abcde", "c", 2},
		{"single char not found", "abcde", "z", -1},
		{"overlapping pattern", "aaa", "aa", 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := kmp(tt.text, tt.substring)
			if result != tt.expected {
				t.Errorf("kmp(%q, %q) = %d, want %d", tt.text, tt.substring, result, tt.expected)
			}
		})
	}
}
