package main

import (
	"bytes"
	"fmt"
	"strings"
	"unsafe"
)

func main() {
	var ff float32
	fmt.Println(unsafe.Sizeof(ff))

	// 1
	s := ""
	for i := 0; i < 10; i++ {
		s += "a"
	}
	fmt.Println(s)

	// 2
	var buffer bytes.Buffer
	for i := 0; i < 10; i++ {
		buffer.WriteString("a")
	}
	fmt.Println(buffer.String())

	// 3
	sl := []string{"a", "a", "a", "a", "a", "a", "a", "a", "a", "a"}
	fmt.Printf(strings.Join(sl, ""))
}
