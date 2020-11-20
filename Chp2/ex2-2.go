package main

import (
	"fmt"
	"genconv"
	"os"
	"strconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		_, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "2-2: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("Value: 2\nIn pounds:\n")
	}
	genconv.CentToFeet(Centimeters(10))
}
