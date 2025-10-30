package main

import (
	"bytes"
	"fmt"
)

func changeSliceSize(ptr *[]byte) {
	new_slice := []byte("some_new_str")
	*ptr = new_slice
}

func main() {
	bSlice := []byte(" \t\n a lone gopher \n\t\r\n")
	fmt.Printf("%s", bytes.TrimSpace(bSlice)) // a lone gopher
	fmt.Printf("%s", bSlice)                  // \t\n a lone gopher \n\t\r\n
	changeSliceSize(&bSlice)
	fmt.Printf("%s", bSlice)
}
