package api

import "testing"

func TestApi(t *testing.T) {
	want := "API Package"
	if got := api(); got != want {
		t.Errorf("api() = %q, want %q", got, want)
	}
}

//-- CHARTMETRIC INTERFACE TESTS --
func TestGetNeighborArtists(t *testing.T) {
	auth := CMAuth{"REPLACE_ME"}
	param := map[string]interface{}{
		"id":     3388,
		"metric": "cm_artist_rank",
		"limit":  10,
		"type":   "",
	}
	call := Call{param, auth.GetNeighborArtists}
	printResults(call.CallOnce())
}
