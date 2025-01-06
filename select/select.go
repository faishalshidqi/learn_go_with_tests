package _select

import (
	"fmt"
	"net/http"
	"time"
)

const tenMilSecTimeout = 10 * time.Millisecond

func Racer(fasturl, slowurl string) (string, error) {
	/*
		elapsedA := measureResponseTime(fasturl)
		elapsedB := measureResponseTime(slowurl)
		if elapsedA < elapsedB {
			return fasturl
		} else {
			return slowurl
		}
	*/
	return ConfigurableRacer(fasturl, slowurl, tenMilSecTimeout)
}

func ConfigurableRacer(fasturl, slowurl string, timeout time.Duration) (string, error) {
	select {
	case <-ping(fasturl):
		return fasturl, nil
	case <-ping(slowurl):
		return slowurl, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("timed out waiting for %s and %s", fasturl, slowurl)
	}
}

func ping(url string) chan struct{} {
	ch := make(chan struct{})
	go func() {
		http.Get(url)
		close(ch)
	}()
	return ch
}

func measureResponseTime(url string) time.Duration {
	start := time.Now()
	http.Get(url)
	elapsed := time.Since(start)
	return elapsed
}
