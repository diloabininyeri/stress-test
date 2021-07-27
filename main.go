package main

import (
	"flag"
	"fmt"
	"stress-test/Request"
	"sync"
	"time"
)

func main() {

	counter := flag.Int("count", 1000, "number of simultaneous requests")
	url := flag.String("url", "http://localhost", "url to test")
	data := flag.String("data", "null", "Specifies the data part to be sent")
	flag.Parse()
	startTest(*counter, *url, *data)
}

func startTest(counter int, url string, data string) {
	var count int = counter
	var uri string = url
	startTime := time.Now()
	var wg sync.WaitGroup
	wg.Add(count)
	for i := 0; i < count; i++ {
		go Request.Send("get", uri, data, &wg)
	}
	wg.Wait()
	fmt.Printf("%d tane işlendi\n", Request.GetCounter())
	endTime := time.Now()
	diffTime := endTime.Sub(startTime)

	fmt.Println(diffTime)

}
