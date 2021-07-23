package main

import (
	"fmt"
	"os"
)

func main (){
	defer func() {
		if v := recover(); v != nil{
			fmt.Println("recovered", v)
		}
	}()
	f, err := os.Create("new_file")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	_, _ = fmt.Fprintln(f, "New line")
}

