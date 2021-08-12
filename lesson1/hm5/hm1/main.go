package main

import (
	"fmt"
	"sync"
)

const count = 100

func main() {
	var (
		x int
		wg = sync.WaitGroup{}
	)
	wg.Add(count)
	for i :=0; i < count; i +=1 {
		go func(){
			x +=1
			fmt.Println(x)
			wg.Done()
		}()
	}
	wg.Wait()
}
