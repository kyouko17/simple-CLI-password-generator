package main

import "testing"

func TestPasswordGenerator(t *testing.T) {
	t.Run("length of password", func(t *testing.T) {
		got := len(PasswordGenerator(16))
		want := 16

		if got != want {
			t.Errorf("expected length: %d, gotten legth: %d", got, want)
		}
	})
}
