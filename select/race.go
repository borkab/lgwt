package race

import (
	"fmt"
	"net/http"
	"time"
)

func Racer(a, b string) (winner string, err error) {
	/*	aDuration := measureResponseTime(a)
		bDuration := measureResponseTime(b)

		if aDuration < bDuration {
			return a
		}

		return b
	*/
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(10 * time.Second):
		return "", fmt.Errorf("timed out, waitind for %s and %s", a, b)
	}
}

/*
	func measureResponseTime(url string) time.Duration {
		start := time.Now()
		http.Get(url)
		return time.Since(start)
	}
*/
func ping(url string) chan struct{} {
	ch := make(chan struct{})
	go func() {
		http.Get(url)
		close(ch)
	}()
	return ch
}
