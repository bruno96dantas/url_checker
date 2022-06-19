package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"runtime"
	"strings"
	"sync"
)

func checkAndSaveBody(url string, wg *sync.WaitGroup) {

	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("%s is DOWN!\n", url)
	} else {

		fmt.Printf("Status Code: %d ", resp.StatusCode)
		if resp.StatusCode == 200 {
			bodyBytes, err := ioutil.ReadAll(resp.Body)

			file := strings.Split(url, "//")[1]
			file += ".txt"

			fmt.Printf("Writing response Body to %s\n", file)
			err = ioutil.WriteFile(file, bodyBytes, 0664)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
	wg.Done()
}
func main() {

	urls := []string{"https://www.golang.org", "https://www.google1.com", "https://www.medium.com"}

	var wg sync.WaitGroup

	wg.Add(len(urls))

	for _, url := range urls {
		go checkAndSaveBody(url, &wg)
	}

	fmt.Println("No. of Goroutines:", runtime.NumGoroutine())

	wg.Wait()
}
