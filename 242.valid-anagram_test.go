package leet

import "testing"

type test struct {
	arg1 string
	arg2 string
	want bool
}

func Test_isAnagram(t *testing.T) {
	tests := []test{
		{"anagram", "nagaram", true},
		{"car", "rat", false},
		{"ac", "bb", false},
	}
	for _, tt := range tests {
		t.Run("isAnagram", func(t *testing.T) {
			got := isAnagram(tt.arg1, tt.arg2)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
