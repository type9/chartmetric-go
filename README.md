# 🔗 chartmetric-go
An abstraction of Charmetric's API in Go with automated concurrency.

[![Go Report Card](https://goreportcard.com/badge/github.com/type9/chartmetric-go)](https://goreportcard.com/report/github.com/type9/chartmetric-go)

### Packages Included

1. [CallMaker](#callmaker)

## CallMaker
See presentation slides on module implementation [GoogleSlides](https://docs.google.com/presentation/d/1MQ9I8GNA6lPCY_egHAE68EqLH_E69H0O1ArUNlA6Jzs/edit?usp=sharing)

Allows for the creation of reusable API call objects. Automates concurrent API calls from a list of url parameters.

- Installation

```bash
go get github.com/type9/chartmetric-go/api
```

- Usage

In order to make a call object you need to specify two things, a request formatting function and parameters

# Formatting function example (invokes parameters, returns *http.Request)
```golang
func getPokemonByName(kwargs map[string]interface{}) *http.Request {
	reqBody, err := json.Marshal(map[string]interface{}{})

	url := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s", kwargs["name"])

	req, err := http.NewRequest("GET", url, bytes.NewBuffer(reqBody)) // creating a custom http request, can be any
	req.Header.Set("Content-type", "application/json")
	if err != nil {
		log.Fatalln(err)
	}

	return req
}
```

# Then, specify literal parameters and the formatting function into Call{} objects
```golang
func TestCallOnce() {
	param := map[string]interface{}{
		"name": "ditto",
	}
	call := Call{param, getPokemonByName} // creates call object with parameters and formatting function
	call.CallOnce() //executes the call using Client.Do() and returns http.Response object
}
```

# For a multiple API calls on a same url with varying parameters, use the MultiCall object
```golang
func TestCallMulti(callback) {
	params := []map[string]interface{}{ // list of parameters
		{"name": "ditto"},
		{"name": "charizard"},
		{"name": "mew"},
		{"name": "pikachu"},
		{"name": "lapras"},
	}
	mcall := MultiCall{params, getPokemonByName} // create MultiCall object with a list of parameters and formatting function
	mcall.CallMulti(callback) // specifiy a callback function that accepts http.Repsonse objects to process the data as you wish
}
```

## Getting Started