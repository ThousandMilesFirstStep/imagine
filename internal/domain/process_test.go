package domain

import "testing"

func TestGetCoordinatesFromPosition(t *testing.T) {
	w1, h1, w2, h2 := 800, 400, 600, 300

	t.Run("is good for TopLeft", func(t *testing.T) {
		x, y := GetCoordinatesFromPosition(w1, h1, w2, h2, TopLeft, 15)

		if x != 15 || y != 15 {
			t.Fail()
		}
	})

	t.Run("is good for TopRight", func(t *testing.T) {
		x, y := GetCoordinatesFromPosition(w1, h1, w2, h2, TopRight, 15)

		if x != 185 || y != 15 {
			t.Fail()
		}
	})

	t.Run("is good for BottomLeft", func(t *testing.T) {
		x, y := GetCoordinatesFromPosition(w1, h1, w2, h2, BottomLeft, 15)

		if x != 15 || y != 85 {
			t.Fail()
		}
	})

	t.Run("is good for BottomRight", func(t *testing.T) {
		x, y := GetCoordinatesFromPosition(w1, h1, w2, h2, BottomRight, 15)

		if x != 185 || y != 85 {
			t.Fail()
		}
	})

	t.Run("is good for Center", func(t *testing.T) {
		x, y := GetCoordinatesFromPosition(w1, h1, w2, h2, Center, 15)

		if x != 100 || y != 50 {
			t.Fail()
		}
	})
}
