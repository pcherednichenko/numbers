package numbers

import (
	"log"
	"sort"
	"sync"
	"time"
)

const (
	timeout = 500

	infoCollectAllNumbersTpl = "Finished collect numbers from %d sources"

	errTooLongTimeoutTpl = "error: url skipped because take too long time, maximum: %d microseconds"
)

type (
	numbers = []int

	results struct {
		mu      sync.Mutex
		results [][]int
	}
)

// CollectNumbersFromURLS collect numbers from urls, group and sort them
func CollectNumbersFromURLS(urls []string) []int {
	var (
		wg     sync.WaitGroup
		r      results
		count  = len(urls)
		result = make(chan numbers, count)
	)
	r.results = make([][]int, count)

	wg.Add(1)
	go func(result chan numbers) {
		defer wg.Done()
		success := 0
		for {
			select {
			case res := <-result:
				r.mu.Lock()
				r.results[success] = res
				r.mu.Unlock()
				success++
			case <-time.After(time.Millisecond * timeout):
				result = nil
				log.Printf(errTooLongTimeoutTpl, timeout)
				return
			}
			if success == count {
				log.Printf(infoCollectAllNumbersTpl, count)
				return
			}
		}
	}(result)
	for _, url := range urls {
		go getNumbersFromURL(url, result)
	}
	wg.Wait()

	collection := make(map[int]int)
	// there will not be a race condition, but if in the future the functionality will expand - this can be a problem
	// so I added the lock here
	r.mu.Lock()
	// remove duplicates and get a slice of all the digits
	for _, numbers := range r.results {
		for _, number := range numbers {
			collection[number] = number
		}
	}
	r.mu.Unlock()

	return sortMap(collection)
}

// sortMap of number into slice
func sortMap(unsorted map[int]int) []int {
	resultNumbers := make([]int, len(unsorted))
	t := 0 // faster because we set len of slice, so we don't need to append (slow operation)
	for _, number := range unsorted {
		resultNumbers[t] = number
		t++
	}
	sort.Ints(resultNumbers)
	return resultNumbers
}
