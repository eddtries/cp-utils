package cputils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSum(t *testing.T) {
	t.Run("test a valid array", func(t *testing.T) {
		want := 10
		got := sum([]int{1, 2, 3, 4})

		assert.Equal(t, want, got)
	})

	t.Run("test an empty array", func(t *testing.T) {
		want := 0
		got := sum([]int{})

		assert.Equal(t, want, got)
	})
}
