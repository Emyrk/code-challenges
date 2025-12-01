package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input := must(os.ReadFile("input.txt"))
	lines := strings.Split(string(input), "\n")
	current := 50
	totalZeroes := 0
	for _, l := range lines {
		dir := l[0]
		mag := must(strconv.Atoi(l[1:]))
		switch dir {
		case 'L':
			current -= mag
		case 'R':
			current += mag
		}

		if current < 0 {
			current = 100 - ((current * -1) % 100)
		}
		if current >= 100 {
			current = current % 100
		}

		if current == 0 {
			totalZeroes++
		}

		fmt.Println(current)
	}
	fmt.Println("Total zeroes:", totalZeroes)
}

func must[V any](value V, err error) V {
	if err != nil {
		panic(err)
	}
	return value
}
