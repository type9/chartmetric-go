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
	auth := CMAuth{"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MjE1NTg4MiwidGltZXN0YW1wIjoxNjE0OTM4NTM0NTUxLCJpYXQiOjE2MTQ5Mzg1MzQsImV4cCI6MTYxNDk0MjEzNH0.RiSKxCXn8FDOlhNlK5j_jerRU1mNobSwAc9CEoTCvmM"}
	param := map[string]interface{}{
		"id":     3388,
		"metric": "cm_artist_rank",
		"limit":  10,
		"type":   "",
	}
	call := Call{param, auth.GetNeighborArtists}
	printResults(call.CallOnce())
}
