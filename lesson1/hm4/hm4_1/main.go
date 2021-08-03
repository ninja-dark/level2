package main

func main(){
	var workers = make(chan struct{}, 1000)

	for i:=1; i <= 1000; i++ {
		workers <- struct{}{}

		go func(job int) {
			defer func() {
				<-workers
			}()
		}
	}
}
