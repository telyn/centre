package main

import (
	"fmt"
	"os"
)

func main() {
	headings := os.Args[1:]
	width := 30

	for _, h := range headings {
		equals := width - (len(h) / 2)
		for i := 0; i < equals; i++ {
			fmt.Print("=")
		}
		fmt.Printf(" %s ", h)
		if len(h)%2 == 0 {
			equals++
		}
		for i := 0; i < equals; i++ {
			fmt.Print("=")
		}
		fmt.Printf("\r\n")
	}

}
