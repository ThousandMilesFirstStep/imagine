package domain

import (
	"testing"
)

func TestNewColorFromHex(t *testing.T) {
	hex := "#ff6432"
	color, err := NewColorFromHex(hex)

	t.Run("run without error", func(t *testing.T) {
		if err != nil {
			t.Fail()
		}
	})

	t.Run("converts red properly", func(t *testing.T) {
		if color.R != 255 {
			t.Errorf("Expect %d to be %d", color.R, 255)
		}
	})

	t.Run("converts green properly", func(t *testing.T) {
		if color.G != 100 {
			t.Errorf("Expect %d to be %d", color.G, 100)
		}
	})

	t.Run("converts blue properly", func(t *testing.T) {
		if color.B != 50 {
			t.Errorf("Expect %d to be %d", color.B, 50)
		}
	})
}
