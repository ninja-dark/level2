package main

import (
	"github.com/pkg/profile"
	"runtime"
	"sync"
)

const count = 1000

func main() {
	defer profile.Start(profile.TraceProfile, profile.ProfilePath(".")).Stop()
	var (
		counter int
		mm      sync.Mutex
		wg      sync.WaitGroup
	)
	wg.Add(count)
	for i := 0; i < count; i += 1 {
		if i == 500 {
			runtime.Gosched()
		}
		go func() {
			defer wg.Done()
			mm.Lock()
			counter += 1
			mm.Unlock()
		}()
	}
	wg.Wait()
}
