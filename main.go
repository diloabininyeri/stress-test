package main

import (
	"fmt"
	"github.com/fatih/color"
	"load-test/Request"
	"sync"
	"time"
)

func signature() {

	color.Green("created by Dılo")
}

func main() {
	signature()
	a()
}

func a() {
	var count int = 100000
	startTime := time.Now()
	var wg sync.WaitGroup
	wg.Add(count)
	for i := 0; i < count; i++ {
		go Request.Send("get", "http://dpdev.dhdestek.com/", "dd", &wg)
	}
	wg.Wait()
	fmt.Printf("%d tane işlendi\n", Request.GetCounter())
	endTime := time.Now()
	diffTime := endTime.Sub(startTime)

	fmt.Println(diffTime)

}
