package main

import (
	"log"
	"net/http"
)

type Site struct {
	URL string
}

type Result struct {
	URL    string
	Status int
}

func crawl(wId int, sites <-chan Site, results chan<- Result) {
	for site := range sites {
		log.Printf("Worker ID %d\n", wId)
		log.Printf("Site URL %s\n", site.URL)

		resp, err := http.Get(site.URL)

		if err != nil {
			log.Println(err.Error())
		}

		results <- Result{
			URL:    site.URL,
			Status: resp.StatusCode,
		}
	}
}

func main() {
	siteChannel := make(chan Site, 2)
	resultChannel := make(chan Result, 2)

	// run 3 workers
	for w := 1; w <= 3; w++ {
		go crawl(w, siteChannel, resultChannel)
	}

	urls := []string{
		"https://example.com",
		"https://google.com",
		"https://liepajniekiem.lv",
	}

	for _, url := range urls {
		siteChannel <- Site{URL: url}
	}
	close(siteChannel)

	for site := 1; site <= len(urls); site++ {
		result := <-resultChannel

		log.Println(result)
	}
	close(resultChannel)

	// for result := range resultChannel { // always waits
	// 	log.Println(result)
	// }
}
