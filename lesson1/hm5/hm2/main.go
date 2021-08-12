package main

import (
	"fmt"
	"sync"
)

func main() {
	var mm sync.Mutex
	mm.Lock()
	foo(&mm)
	mm.Lock()
	foo(&mm)
}

func foo(mm1 *sync.Mutex) {
	defer mm1.Unlock()
	fmt.Println("well done")

}
