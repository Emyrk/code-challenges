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
		for i := 0; i < mag; i++ {
			current = rotate(current, dir == 'R')
			if current == 0 {
				totalZeroes++
			}
		}
	}
	fmt.Println("Total zeroes:", totalZeroes)
}

func rotate(current int, dir bool) int {
	if dir {
		current++
	} else {
		current--
	}

	if current == -1 {
		return 99
	}
	if current == 100 {
		return 0
	}
	return current
}

func must[V any](value V, err error) V {
	if err != nil {
		panic(err)
	}
	return value
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
