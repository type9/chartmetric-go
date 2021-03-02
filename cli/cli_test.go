package cli

import "testing"

func TestCli(t *testing.T) {
	want := "CLI Package"
	if got := cli(); got != want {
		t.Errorf("cli() = %q, want %q", got, want)
	}
}
