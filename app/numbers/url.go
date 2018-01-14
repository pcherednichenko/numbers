package numbers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	errWhileSendRequestTpl       = "error: while send request: %v"
	errURLReturnBadStatusCodeTpl = "error: url %s return bad status code %d"
	errUnmarshalBodyTpl          = "error: while unmarshal body %v"
)

type response struct {
	Numbers []int `json:"numbers"`
}

// getNumbersFromURL send request and parse response
func getNumbersFromURL(url string, result chan numbers) {
	resp, err := http.Get(url)
	if err != nil {
		log.Printf(errWhileSendRequestTpl, err)
		writeToChannel([]int{}, result)
		return
	}
	if resp.StatusCode >= http.StatusBadRequest {
		log.Printf(errURLReturnBadStatusCodeTpl, url, resp.StatusCode)
		writeToChannel([]int{}, result)
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf(errURLReturnBadStatusCodeTpl, url, resp.StatusCode)
		writeToChannel([]int{}, result)
		return
	}
	var respBody response
	err = json.Unmarshal(body, &respBody)
	if err != nil {
		log.Printf(errUnmarshalBodyTpl, err)
		writeToChannel([]int{}, result)
		return
	}

	writeToChannel(respBody.Numbers, result)
}

// writeToChannel numbers result
func writeToChannel(numbers []int, result chan numbers) {
	// maybe channel are closed because of timeout
	if result != nil {
		result <- numbers
	}
}
