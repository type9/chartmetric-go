package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"testing"
)

//invokes http response, decodes into json and writes into file with specified name
func writeResults(resp *http.Response, fileName string) {
	fName := fileName
	file, err := os.Create(fName)
	if err != nil {
		log.Fatalf("Cannot create file %q: %s\n", fName, err)
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	defer file.Close()
	defer resp.Body.Close()

	reader := strings.NewReader(string(body))

	dec := json.NewDecoder(reader)
	enc := json.NewEncoder(file)
	enc.SetIndent("", "  ")

	objects := []map[string]interface{}{}
	for {
		var m map[string]interface{}
		if err := dec.Decode(&m); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}

		objects = append(objects, m)
	}
	enc.Encode(objects)
}

func printResults(resp *http.Response) {
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(string(body))
}

//example formatter function for Pokemon API
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

func TestCallOnce(t *testing.T) {
	param := map[string]interface{}{
		"name": "ditto",
	}
	call := Call{param, getPokemonByName}
	writeResults(call.CallOnce(), "./output/calloncetest_results.json")
}

var iter = 0

func respIterate(resp *http.Response) {
	fname := fmt.Sprintf("./output/multicalltest_result_%d.json", iter)
	writeResults(resp, fname)
	iter++
}

//Tests concurrent calls using MultiCall object
func TestCallMulti(t *testing.T) {
	params := []map[string]interface{}{
		{"name": "ditto"},
		{"name": "charizard"},
		{"name": "mew"},
		{"name": "pikachu"},
		{"name": "lapras"},
	}
	mcall := MultiCall{params, getPokemonByName}
	mcall.CallMulti(respIterate)
	iter = 0
}

//Tests a non-concurrent method of getting multiple calls worth of data
func TestCallMultiIterated(t *testing.T) {
	params := []map[string]interface{}{
		{"name": "ditto"},
		{"name": "charizard"},
		{"name": "mew"},
		{"name": "pikachu"},
		{"name": "lapras"},
	}
	for _, param := range params {
		call := Call{param, getPokemonByName}
		respIterate(call.CallOnce())
	}
	iter = 0
}
