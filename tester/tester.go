package tester

import (
	"net/http"
	"sync"
	"time"
)

type TestResult struct {
	TotalRequests    int
	FailedRequests   int
	SuccessCount     int
	StatusCodeCounts map[int]int
	Duration         time.Duration
}

func RunLoadTest(url string, totalRequests, concurrency int) TestResult {
	var wg sync.WaitGroup
	result := TestResult{
		TotalRequests:    totalRequests,
		StatusCodeCounts: make(map[int]int),
	}

	requests := make(chan int, totalRequests)
	start := time.Now()

	for i := 0; i < totalRequests; i++ {
		requests <- i
	}
	close(requests)

	mutex := &sync.Mutex{}

	for i := 0; i < concurrency; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for range requests {
				resp, err := http.Get(url)
				if err != nil {
					mutex.Lock()
					result.StatusCodeCounts[0]++
					mutex.Unlock()
					continue
				}
				mutex.Lock()
				result.StatusCodeCounts[resp.StatusCode]++
				if resp.StatusCode == 200 {
					result.SuccessCount++
				}
				mutex.Unlock()
				resp.Body.Close()
			}
		}()
	}

	wg.Wait()
	result.Duration = time.Since(start)
	result.FailedRequests = result.TotalRequests - result.SuccessCount
	return result
}
