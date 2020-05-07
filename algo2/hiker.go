package main

import (
	"fmt"
	"strconv"
	"strings"
)

func isAnumber(param string) (int, error) {
	result, err := strconv.Atoi(param)
	if err != nil {
		return 0, err
	}
	return result, nil
}

func main() {
	var (
		countInput  string
		valleyInput string
		valleys     []string
		count       int
	)
	fmt.Print("Input the number of steps : ")
	for {
		fmt.Scan(&countInput)
		result, err := isAnumber(countInput)
		if err != nil {
			fmt.Printf("%s not a number, please input number : ", err)
			continue
		}
		count = result
		break
	}

	fmt.Printf("\nplease Input the steps U or D (use white space) : ")
	for len(valleys) < count {
		fmt.Scan(&valleyInput)
		valleys = append(valleys, strings.Split(valleyInput, " ")...)
	}

	fmt.Printf("\ntotal of journey are : %d\n", countingValley(valleys))
}

func countingValley(socks []string) int {
	m := make(map[string]int)
	for _, v := range socks {
		m[v]++
	}

	return abs(m["U"] - m["D"])
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
