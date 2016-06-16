package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"sync"
)

func getPage(address string) (string, error) {
	log.Printf("logading %s", address)
	resp, err := http.Get(address)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	log.Printf("%s loaded", address)
	return string(body), err
}

// start OMIT
func main() {
	pages := []string{"http://tsuru.io", "http://globo.com", "http://g1.globo.com"}
	var wg sync.WaitGroup
	for _, page := range pages {
		wg.Add(1)
		go func(page string) {
			defer wg.Done()
			getPage(page)
		}(page)
	}
	wg.Wait()
}

// end OMIT
