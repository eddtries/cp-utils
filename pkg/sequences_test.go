package cputils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCollatzSequence(t *testing.T) {
	t.Run("test valid collatz sequence", func(t *testing.T) {
		want := []int{13, 40, 20, 10, 5, 16, 8, 4, 2, 1}
		got, _ := CollatzSequence(13)

		assert.ElementsMatch(t, want, got)
	})

	t.Run("test collatz sequence of 1", func(t *testing.T) {
		want := []int{1}
		got, _ := CollatzSequence(1)

		assert.ElementsMatch(t, want, got)
	})

	t.Run("test collatz sequence with negative integer", func(t *testing.T) {
		_, err := CollatzSequence(-10)
		assert.Error(t, err)
	})
}
