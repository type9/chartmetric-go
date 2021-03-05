package api

import (
	"log"
	"net/http"
	"sync"
	"time"
)

const timeoutDuration = 20 * time.Second

type reqFormatter func(map[string]interface{}) *http.Request // function that takes an array of paratmeters to formats it into a url string

//Call object takes an array of parameters and a function which formats it into a url
type Call struct {
	Param  map[string]interface{}
	Format reqFormatter
}

//MultiCall takes an 2d array of parameters and a function which formats it into a url
type MultiCall struct {
	Params []map[string]interface{}
	Format reqFormatter
}

//CallOnce takes a single Call object and forms a full url and fetches the response
func (call Call) CallOnce() *http.Response {
	timeout := time.Duration(timeoutDuration)
	client := http.Client{
		Timeout: timeout,
	}

	req := call.Format(call.Param)

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	return resp
}

//async version of CallOnce that invokes a call method and utilizes a channel and a waitgroup
func (call Call) asyncCall(ch chan<- *http.Response, wg *sync.WaitGroup) {
	defer wg.Done()
	timeout := time.Duration(timeoutDuration)
	client := http.Client{
		Timeout: timeout,
	}

	req := call.Format(call.Param)

	resp, err := client.Do(req)
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
		go func(param map[string]interface{}) {
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
