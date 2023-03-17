package main

import (
	"fmt"
	"net/http"
	"sync"
)

var urls = []string{
	"https://google.com",
	"https://twitter.com",
	"https://tutorialedge.net",
}

func crawl(url string, waitGroup *sync.WaitGroup) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(resp.Status)

	waitGroup.Done()
}

var globalMutex sync.Mutex

func synchronizedFoo(mutex *sync.Mutex) {
	mutex.Lock()

	mutex.Unlock()
}

func main() {
	var waitGroup sync.WaitGroup

	for _, url := range urls {
		waitGroup.Add(1)
		go crawl(url, &waitGroup)
	}

	waitGroup.Wait()

	synchronizedFoo(&globalMutex)
}
