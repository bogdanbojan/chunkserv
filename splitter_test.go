package main

import (
	"testing"
)

func TestSplitLine(t *testing.T) {
	assertEqual := func(t testing.TB, got Lines, want Lines) {
		t.Helper()
		for i, v := range got {
			if v != want[i] {
				t.Errorf("got %q want %q", got, want)
			}
		}
	}

	t.Run("same line", func(t *testing.T) {
		testLine := "1|Mavra|Malec|mmalec0@usa.gov|Female|229.215.245.102"
		got := splitLine(testLine)
		want := Lines{"1", "Mavra", "Malec", "mmalec0@usa.gov", "Female", "229.215.245.102"}

		assertEqual(t, got, want)
	})
	//if !isEqual(got, want) {
	//	t.Errorf("got %q want %q", got, want)
	//}

	//if got != want {
	//	t.Errorf("got %q want %q", got, want)
	//}

}

func isEqual(line1 Lines, line2 Lines) bool {
	for i, v := range line1 {
		if v != line2[i] {
			return false
		}
	}
	return true
}

func TestIsValid(t *testing.T) {}
