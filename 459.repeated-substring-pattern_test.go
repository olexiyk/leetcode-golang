package leet

import "testing"

func Test_repeated(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{"abab", true},
		{"car", false},
		{"aaaa", true},
		{"abcabcabcabc", true},
	}
	for _, tt := range tests {
		t.Run("isAnagram", func(t *testing.T) {
			got := repeatedSubstringPattern(tt.input)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
