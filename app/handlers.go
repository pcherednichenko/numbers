package app

import (
	"log"
	"net/http"

	"encoding/json"

	"github.com/pcherednichenko/numbers/app/numbers"
)

const (
	urlParam = "u"

	errCantFindParam            = "error: impossible to find '%s' parameter in the query"
	errWhileMarshallResponseTpl = "error: while marshall response: %v"
)

type response struct {
	Numbers []int `json:"numbers"`
}

func numbersHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// 200 because of task description, in real life better return bad status if err
	w.WriteHeader(http.StatusOK)
	params := r.URL.Query()
	urls, ok := params[urlParam]
	if !ok {
		log.Printf(errCantFindParam, urlParam)
		writeEmptyResponse(w)
		return
	}
	var resp response
	resp.Numbers = numbers.CollectNumbersFromURLS(urls)
	resultResponse, err := json.Marshal(resp)
	if err != nil {
		log.Printf(errWhileMarshallResponseTpl, err)
		writeEmptyResponse(w)
		return
	}
	w.Write(resultResponse)
}

func writeEmptyResponse(w http.ResponseWriter) {
	var resp response
	resp.Numbers = []int{} // because we need "numbers": [] not "numbers": null in json
	emptyResp, err := json.Marshal(resp)
	if err != nil {
		log.Printf(errWhileMarshallResponseTpl, err)
		return
	}
	w.Write(emptyResp)
}
