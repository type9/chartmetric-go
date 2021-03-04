package api

import (
	"log"
	"net/http"
	"sync"
)

type urlFormatter func([]string) string // Function that takes an array of paratmeters to formats it into a url string

//Call object takes an array of parameters and a function which formats it into a url
type Call struct {
	Param  []string
	Format urlFormatter
}

//MultiCall takes an 2d array of parameters and a function which formats it into a url
type MultiCall struct {
	Params [][]string
	Format urlFormatter
}

//CallOnce takes a single Call object and forms a full url and fetches the response
func (call Call) CallOnce() *http.Response {
	url := call.Format(call.Param)

	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}

	return resp
}

//helper call function for multicall
func (call Call) asyncCall(ch chan<- *http.Response, wg *sync.WaitGroup) {
	defer wg.Done()
	url := call.Format(call.Param)

	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}

	ch <- resp
}

//CallMulti takes a MultiCall object and attempts to concurrently fetch the url for each set of parameters
func (mcall MultiCall) CallMulti(callback func(*http.Response)) {
	respChan := make(chan *http.Response)

	wg := sync.WaitGroup{}
	for _, param := range mcall.Params { //for each set of parameters
		wg.Add(1)

		//capture parameters and run an async api call
		go func(param []string) {
			call := Call{param, mcall.Format}
			call.asyncCall(respChan, &wg)
		}(param)

	}

	go func() { //async detect when workers are done and close the response channel
		wg.Wait()
		close(respChan)
	}()

	for res := range respChan {
		callback(res)
	}
}
