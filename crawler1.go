package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

// start OMIT
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

func main() {
	pages := []string{"http://tsuru.io", "http://globo.com", "http://g1.globo.com"}
	for _, page := range pages {
		_, err := getPage(page)
		if err != nil {
			log.Printf("Error getting %s content", page)
		}
	}
}

// end OMIT
