package main

import (
	"fmt"
	"time"
)

var printValues bool
var terminateProcess uint64

func startPrintingValues(pv chan bool) {
	for {
		pv <- printValues
	}
}

func endProcess(tp chan uint64) {
	for {
		tp <- terminateProcess
	}
}

func process(id uint64, printValues chan bool, terminateProcess chan uint64) {
	i := uint64(0)
	for {
		pv := <-printValues
		tp := <-terminateProcess
		// print values
		if pv {
			fmt.Printf("id %d: %d\n", id, i)
		}

		i = i + 1
		time.Sleep(time.Millisecond * 500)
		// terminate process
		if tp == id {
			break
		}
	}
}

func main() {
	var opt string
	var totalProcesses uint64
	// channels
	var pv chan bool = make(chan bool)
	var tp chan uint64 = make(chan uint64)

	// innit channels
	go startPrintingValues(pv)
	go endProcess(tp)

	for opt != "0" {
		fmt.Println("1-. Add new process")
		fmt.Println("2.- Show processes info")
		fmt.Println("3.- Delete process")
		fmt.Println("0.- Exit")
		fmt.Print("\nSelect an option: ")
		fmt.Scanln(&opt)

		switch opt {
		case "1":
			go process(totalProcesses+1, pv, tp)
			fmt.Println("Added process", totalProcesses+1)
			totalProcesses++
			break
		case "2":
			//pv <- true
			printValues = true
			fmt.Scanln()
			printValues = false
		}
	}
}
