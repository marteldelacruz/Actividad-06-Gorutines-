package main

import (
	"fmt"
	"time"
)

var printValues bool

func Proceso(id uint64) {
	i := uint64(0)
	for {
		if printValues {
			fmt.Printf("id %d: %d", id, i)
		}
		i = i + 1
		time.Sleep(time.Millisecond * 500)
	}
}

func main() {
	var opt string

	// main menu
	for opt != "0" {
		fmt.Println("1-. Add new process")
		fmt.Println("2.- Show processes info")
		fmt.Println("3.- Delete process")
		fmt.Println("0.- Exit")
		fmt.Print("\nSelect an option: ")

		fmt.Scanln(&opt)
	}
}
