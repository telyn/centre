package main

import (
	"flag"
	"fmt"
	"os"
)

func centre(str, padding string, width uint) (out string) {
	if len(str) >= int(width) {
		return str
	}
	padwidth := (int(width) - 2 - len(str)) / 2
	for i := 0; i < padwidth; i += len(padding) {
		out += padding
	}
	out += fmt.Sprintf(" %s ", str)
	if len(str)%2 == 1 {
		padwidth++
	}
	for i := 0; i < padwidth; i += len(padding) {
		out += padding
	}
	return
}

func main() {
	headings := os.Args[1:]
	width := flag.Uint("width", 30, "The width you want your headings padded to")
	padding := flag.String("padding", "=", "The string to use as padding.")

	for _, h := range headings {
		fmt.Println(centre(h, *padding, *width))
	}

}
