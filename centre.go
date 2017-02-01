package main

import (
	"flag"
	"fmt"
	"os"
)

func centre(str, padding string, width uint) string {
	text := []rune(str)
	if len(text) >= int(width) {
		return str
	}
	padrunes := []rune(padding)

	out := make([]rune, int(width))
	pos := 0

	padwidth := int(width)/2 - 2 - len(text)/2
	if len(str)%2 == 0 {
		padwidth++
	}
	for i := 0; i < padwidth; i++ {
		out[pos] = padrunes[i%len(padrunes)]
		pos++
	}
	out[pos] = ' '
	pos++
	for i := 0; i < len(text); i++ {
		out[pos] = text[i]
		pos++
	}
	out[pos] = ' '
	pos++

	// if padwidth * 2 + 2 + len(str) < width {
	if len(str)%2 == 1 {
		padwidth++
	}
	for i := 0; i < padwidth; i++ {
		out[pos] = padrunes[i%len(padrunes)]
		pos++
	}
	return string(out)
}

func main() {
	headings := os.Args[1:]
	width := flag.Uint("width", 30, "The width you want your headings padded to")
	padding := flag.String("padding", "=", "The string to use as padding.")

	for _, h := range headings {
		fmt.Println(centre(h, *padding, *width))
	}

}
