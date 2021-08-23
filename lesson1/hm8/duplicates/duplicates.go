package duplicates

import (
	"fmt"
	"io"
)

type Duplicate struct {
	Value1 string
	Value2 string
}

type DuplicateHandler func(Duplicate)

func GetWriter(writer io.Writer) DuplicateHandler {
	return func(d Duplicate) {
		fmt.Fprintf(writer, d.Value2)
	}
}

func GetCSVWriter(writer io.Writer) DuplicateHandler {
	return func(d Duplicate) {
		fmt.Fprintf(writer, "\" %s\",\"%s\"\n", d.Value1, d.Value2)
	}
}

func ApplyFuncToChan(duplicates <-chan Duplicate, handler DuplicateHandler) {
	for d := range duplicates {
		handler(d)
	}
}