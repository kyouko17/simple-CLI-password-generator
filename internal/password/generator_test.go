package password

import "testing"

func TestPasswordGenerator(t *testing.T) {
	t.Run("length of password", func(t *testing.T) {
		password, err := Generate(16)
		got := len(password)
		want := 16

		if err != nil {
			t.Errorf("unexpected error: %g", err)
		}
		if got != want {
			t.Errorf("expected length: %d, gotten legth: %d", got, want)
		}
	})
	t.Run("length out of range", func(t *testing.T) {
		password, err := Generate(3)
		got := password
		want := ""

		if err != ErrNotAcceptableLength {
			t.Errorf("")
		}
		if got != want {
			t.Errorf("got: %s, want: %s", got, want)
		}
	})
}
