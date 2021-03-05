package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

//CMAuth requires an access key which can be retrieved using a ChartMetric refresh token
type CMAuth struct {
	Key string
}

func (auth CMAuth) formatReq(req *http.Request) { //TODO: Import refresh token using viper and gain valid refresh token
	req.Header.Set("Authorization", "Bearer "+auth.Key)
}

//Formatters reutrn a formatted request using kwargs and authentication object

//GetNeighborArtists has parameters (* means required ) => *id, metric (string), limit (int), type (string)
func (auth CMAuth) GetNeighborArtists(kwargs map[string]interface{}) *http.Request {
	reqBody, err := json.Marshal(map[string]interface{}{})
	if err != nil {
		log.Fatalln(err)
	}

	//TODO: Fix URL formatting such that optional parameters are ommitted if null
	url := fmt.Sprintf("https://api.chartmetric.com/api/artist/%d/neighboring-artists?metric=%s&limit=%d", kwargs["id"], kwargs["metric"], kwargs["limit"])

	req, err := http.NewRequest("GET", url, bytes.NewBuffer(reqBody))
	auth.formatReq(req) //formats header
	if err != nil {
		log.Fatalln(err)
	}
	return req
}
