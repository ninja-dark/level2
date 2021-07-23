package main

import (
	"errors"
	"fmt"
	"time"
)
type MyError struct {
	time string
	text string
}

func New(text string) error{
	return &MyError{
		time: timeError(),
		text: text,
	}
}

func(e *MyError) Error() string {
	return fmt.Sprintf("Time: %s\n Text:\n%s", e.time, e.text)
}

func timeError() string {
	time := time.Now()
	return time.Format( "15:04:05")
}

func main () {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("panic occurred:", err)
		}
	}()
	var a, b, c uint32
	var err error
	err = errors.New("my error")
	fmt.Println(err)

	err = New("my error")
	fmt.Println(err)
	a = 12
	c = a/b
	fmt.Println("Равно: %d", c)

}
