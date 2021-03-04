package api

import (
	"io/ioutil"
	"log"
	"net/http"
	"testing"
)

func printResults(resp *http.Response) {
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(string(body))
}

func pokeURL(param []string) string {
	url := "https://pokeapi.co/api/v2/pokemon/" + param[0]
	return url
}

func TestCallOnce(t *testing.T) {
	param := []string{"ditto"}
	call := Call{param, pokeURL}
	printResults(call.CallOnce())
}

func TestCallMulti(t *testing.T) {
	params := [][]string{
		{"ditto"},
		{"mew"},
		{"charizard"},
		{"pikachu"},
	}
	mcall := MultiCall{params, pokeURL}
	mcall.CallMulti(printResults)
}
