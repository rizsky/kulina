package main

import (
	"fmt"
	"os"
	"regexp"
)

/*
There is an input number: 1.345.679
Write pseudo code (use GoLang) that produces following output:
1000000
300000
40000
5000
600
70
9
*/

func main() {
	var input = "1.345.679"
	reg, err := regexp.Compile(`[^a-zA-Z0-9]+`)
	if err != nil {
		fmt.Println("something went wrong with the input")
		os.Exit(2)
	}

	data := reg.ReplaceAllString(input, "")
	length := len(data)
	zero := "0"

	for i := 0; i < len(data); i++ {
		length--
		if length == 1 {
			zero = ""
		}
		fmt.Printf("%s%0*s\n", string(data[i]), length, zero)
	}
}
