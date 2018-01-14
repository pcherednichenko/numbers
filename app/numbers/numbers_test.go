package numbers

import (
	"testing"
)

// TestCollectNumbers basic test of two worked servers with normal numbers
func TestCollectNumbers(t *testing.T) {
	testServerFirst := testServer(response{Numbers: []int{1, 8, 9}})
	testServerSecond := testServer(response{Numbers: []int{3, 4, 5}})
	urls := []string{testServerFirst.URL, testServerSecond.URL}
	result := CollectNumbersFromURLS(urls)

	if !equalSlice([]int{1, 3, 4, 5, 8, 9}, result) {
		t.Fail()
	}
}

// TestCombiningRepeatedNumbers check that same numbers will be only one time
func TestCombiningRepeatedNumbers(t *testing.T) {
	testServerFirst := testServer(response{Numbers: []int{1, 2, 9}})
	testServerSecond := testServer(response{Numbers: []int{2, 2, 4}})
	urls := []string{testServerFirst.URL, testServerSecond.URL}
	result := CollectNumbersFromURLS(urls)

	if !equalSlice([]int{1, 2, 4, 9}, result) {
		t.Fail()
	}
}

// TestOneServerReturnsEmptyList check if one server return empty list - we get only numbers from other server
func TestOneServerReturnsEmptyList(t *testing.T) {
	testServerFirst := testServer(response{Numbers: []int{1, 9, 2, 1}})
	testServerSecond := testServer(response{Numbers: []int{}})
	urls := []string{testServerFirst.URL, testServerSecond.URL}
	result := CollectNumbersFromURLS(urls)

	if !equalSlice([]int{1, 2, 9}, result) {
		t.Fail()
	}
}

// TestOneServerIsVerySlow testing delay for response of service, here delay is too big, so we skip server
func TestOneServerIsVerySlow(t *testing.T) {
	testServerFirst := testServer(response{Numbers: []int{7, 8, 9}})
	// too long response time for use it
	testServerSecond := testServerWithDelay(response{Numbers: []int{1, 2, 3}}, 700)
	urls := []string{testServerFirst.URL, testServerSecond.URL}
	result := CollectNumbersFromURLS(urls)

	if !equalSlice([]int{7, 8, 9}, result) {
		t.Fail()
	}
}

// TestOneWrongURLAndNormal test that with one wrong url we still get numbers from another
func TestOneWrongURLAndNormal(t *testing.T) {
	testServerFirst := testServer(response{Numbers: []int{20, 21, 23}})
	urls := []string{testServerFirst.URL, "http wrong url com"}
	result := CollectNumbersFromURLS(urls)

	if !equalSlice([]int{20, 21, 23}, result) {
		t.Fail()
	}
}

// equalSlice just compare two slices, return true if equals
func equalSlice(expected []int, actial []int) bool {
	for i, a := range actial {
		if expected[i] != a {
			return false
		}
	}
	return true
}
