package main

import (
	"fmt"
	"net/http"
	"time"
)

type Site struct {
	Url string
}

type Result struct {
	StatusCode int
	Url        string
}

func Job(id int, job chan Site, res chan Result) {
	for i := range job {
		z, _ := http.Get(i.Url)
		res <- Result{StatusCode: z.StatusCode, Url: i.Url}
	}
}
func main() {
	jobs := make(chan Site, 3)
	result := make(chan Result, 3)

	start := time.Now()

	for i := 0; i < 2; i++ {
		go Job(i, jobs, result)
	}

	urls := []string{
		"http://google.com",
		"http://codeforces.com",
		"http://example.com",
		"http://gmail.com",
	}

	for _, url := range urls {
		jobs <- Site{Url: url}
	}

	close(jobs)

	for i := 0; i < 4; i++ {
		res := <-result
		fmt.Println("Url : ", res.Url, " Status : ", res.StatusCode)
	}

	dur := time.Since(start)

	fmt.Println("Time taken : ", dur)
}
