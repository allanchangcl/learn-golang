package main

import (
	"log"
	"net/http"
	"time"
)

type result struct {
	url     string
	err     error
	latency time.Duration
}

func get(url string, ch chan<- result) {
	start := time.Now()

	// use get
	if resp, err := http.Get(url); err != nil {
		// if error, send the results
		// ch <- result{}
		ch <- result{url, err, 0}
	} else {
		t := time.Since(start).Round(time.Microsecond)

		// if successful, also send the results
		// ch <- result{}
		ch <- result{url, nil, t}

		// if no error, make sure sockets get closed
		defer resp.Body.Close()

	}
}

func main() {
	list := []string{
		"https://cheeloy.com",
		"https://amazon.com",
		"https://youtube.com",
		"https://threadykeycaps.com",
		"https://capsphere.com",
	}

	// need make channel for results
	results := make(chan result)

	// need loop over list
	for _, url := range list {
		// starting 4 go routines
		go get(url, results)
	}

	// can only close channels once
	for range list {
		r := <-results

		if r.err != nil {
			log.Printf("%-20s %s\n", r.url, r.err)
		} else {
			log.Printf("%-20s %s\n", r.url, r.latency)
		}
	}
}
