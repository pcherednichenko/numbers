package numbers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

// TestNormalURLWithNormalData check main normal work with normal data
func TestNormalURLWithNormalData(t *testing.T) {
	expected := []int{1, 3, 2}
	testServer := testServer(response{Numbers: expected})
	defer testServer.Close()
	result := make(chan numbers, 1)
	getNumbersFromURL(testServer.URL, result)
	r := <-result
	for i, e := range expected {
		if r[i] != e {
			t.Fail()
		}
	}
}

// TestWrongURL check that with wrong data we not receive any numbers
func TestWrongURL(t *testing.T) {
	result := make(chan numbers, 1)
	getNumbersFromURL("wrong url", result)
	r := <-result
	if len(r) > 0 {
		t.Fail()
	}
}

// TestWrongData check that with wrong data we still get empty slice of numbers
func TestWrongData(t *testing.T) {
	var wrongResp struct {
		Numbers []string `json:"numbers"`
	}
	wrongResp.Numbers = []string{"a", "b", "c"}
	testServer := testServer(wrongResp)
	defer testServer.Close()
	result := make(chan numbers, 1)
	getNumbersFromURL(testServer.URL, result)
	r := <-result
	if len(r) > 0 {
		t.Fail()
	}
}

// TestWrongData check that we can't get numbers if we get wrong status code
func TestNotFoundError(t *testing.T) {
	var errorResp struct {
		Error   string `json:"error"`
		Numbers []int  `json:"numbers"`
	}
	errorResp.Error = "404 Not found error!"
	errorResp.Numbers = []int{1, 2, 3}

	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		resp, err := json.Marshal(errorResp)
		if err != nil {
			t.Fatal(err)
		}
		w.Write(resp)
	}))
	defer testServer.Close()
	result := make(chan numbers, 1)
	getNumbersFromURL(testServer.URL, result)
	r := <-result
	if len(r) > 0 {
		t.Fail()
	}
}

// testServer just provides httptest server for tests
func testServer(response interface{}) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		writeResponse(w, response)
	}))
}

// testServer just provides httptest server for tests with sleep time
func testServerWithDelay(response interface{}, delay int) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if delay != 0 {
			time.Sleep(time.Duration(delay) * time.Millisecond)
		}
		writeResponse(w, response)
	}))
}

// writeResponse write OK response with marshal struct
func writeResponse(w http.ResponseWriter, response interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	resp, _ := json.Marshal(response)
	w.Write(resp)
}
