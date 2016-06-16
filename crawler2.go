package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

// start OMIT
func getPage(address string, c chan string) {
	log.Printf("logading %s", address)
	resp, err := http.Get(address)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	_, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	log.Printf("%s loaded", address)
	c <- address
}

func main() {
	c := make(chan string)
	pages := []string{"http://tsuru.io", "http://globo.com", "http://g1.globo.com"}
	for _, page := range pages {
		go getPage(page, c)
	}
	for url := range c {
		log.Println(url)
	}
}

// end OMIT
