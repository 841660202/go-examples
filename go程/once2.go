package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var once sync.Once

	onceBody := func() {
		fmt.Println("Only once")
	}

	done := make(chan bool)

	for i := 0; i < 10; i++ {
		go func() {
			once.Do(onceBody)
			time.Sleep(time.Second)
			done <- true
		}()
	}

	for i := 0; i < 10; i++ {
		<-done
		fmt.Println("I", i)
	}
}
