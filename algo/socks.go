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
		countInput string
		sockInput  string
		socks      []string
		count      int
	)
	fmt.Print("Input the number of socks : ")
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

	fmt.Printf("\nplease Input the sock (use white space) : ")
	for len(socks) < count {
		fmt.Scan(&sockInput)
		socks = append(socks, strings.Split(sockInput, " ")...)
	}

	fmt.Printf("total of the pair possible are : %d\n", findPair(socks))

}

func findPair(socks []string) int {
	result := 0
	m := make(map[string]int)
	for _, v := range socks {
		m[v]++
		if m[v] == 2 {
			m[v] = 0
			result++
		}
	}

	return result
}
