package leet

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_MoveZeros(t *testing.T) {
	tests := []struct {
		input []int
		want  []int
	}{
		{[]int{1, 2}, []int{1, 2}},
		{[]int{0, 1, 2}, []int{1, 2, 0}},
		{[]int{1, 0, 0, 2}, []int{1, 2, 0, 0}},
	}
	for _, tt := range tests {
		t.Run("isAnagram", func(t *testing.T) {
			moveZeroes(tt.input)
			assert.Equal(t, tt.input, tt.want)
		})
	}
}
