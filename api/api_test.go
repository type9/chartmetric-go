package api

import "testing"

func TestApi(t *testing.T) {
	want := "API Package"
	if got := api(); got != want {
		t.Errorf("api() = %q, want %q", got, want)
	}
}
