package main

import (
	"fmt"
)

func main() {
	slice := make([]int, 100, 100)

	fmt.Println(slice)

	for i := range slice {
		slice[i] = i + 1
	}

	fmt.Println(slice)

	fmt.Println(slice[:10])
	fmt.Println(slice[90:])

	slice = append(slice[:10], slice[90:]...)

	fmt.Println(slice)
	fmt.Println(len(slice))

	for i := range slice {
		if i*2 == len(slice) {
			break
		}
		slice[i], slice[len(slice)-i-1] = slice[len(slice)-i-1], slice[i]

	}
	fmt.Println(slice)

}
