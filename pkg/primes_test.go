package cputils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPrime(t *testing.T) {
	t.Run("13 is prime", func(t *testing.T) {
		want := true
		got := IsPrime(13)

		assert.Equal(t, want, got)
	})

	t.Run("12 is not prime", func(t *testing.T) {
		want := false
		got := IsPrime(12)

		assert.Equal(t, want, got)
	})

	t.Run("negative numbers are not prime", func(t *testing.T) {
		want := false
		got := IsPrime(-4)

		assert.Equal(t, want, got)
	})

	t.Run("1 is not prime", func(t *testing.T) {
		want := false
		got := IsPrime(1)

		assert.Equal(t, want, got)
	})

	t.Run("2 is prime", func(t *testing.T) {
		want := true
		got := IsPrime(2)

		assert.Equal(t, want, got)
	})
}
