package main

import (
	"fmt"
	"sync"
	"time"
)

func count(thing string) {
	for i := 1; i <= 5; i++ {
		fmt.Println(i, thing)
		time.Sleep(time.Millisecond * 500)
	}
}
func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		count("sheep")
		// count("fish")
		wg.Done()
	}()
	wg.Wait()
}
