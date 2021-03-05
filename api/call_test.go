package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func printResults(resp *http.Response) {
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(string(body))
}

func getPokemonByName(kwargs map[string]interface{}) *http.Request {
	reqBody, err := json.Marshal(map[string]interface{}{})
	url := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s", kwargs["name"])
	req, err := http.NewRequest("GET", url, bytes.NewBuffer(reqBody))
	req.Header.Set("Content-type", "application/json")
	if err != nil {
		log.Fatalln(err)
	}
	return req
}

// func TestCallOnce(t *testing.T) {
// 	param := map[string]interface{}{
// 		"name": "ditto",
// 	}
// 	call := Call{param, getPokemonByName}
// 	printResults(call.CallOnce())
// }

// func TestCallMulti(t *testing.T) {
// 	params := []map[string]interface{}{
// 		{"name": "ditto"},
// 		{"name": "charizard"},
// 		{"name": "mew"},
// 		{"name": "pikachu"},
// 		{"name": "lapras"},
// 	}
// 	mcall := MultiCall{params, getPokemonByName}
// 	mcall.CallMulti(printResults)
// }
