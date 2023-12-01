package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	var res int
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		var cur int
		sS := true
		eS := true

		for i, c := range scanner.Text() {
			k := len(scanner.Text()) - 1 - i
			if parsed, err := strconv.Atoi(string(c)); err == nil && sS {
				cur += parsed * 10
				sS = false
			}

			if parsed, err := strconv.Atoi(string(scanner.Text()[k])); err == nil && eS {
				cur += parsed
				eS = false
			}
		}
		res += cur
	}

	fmt.Println(res)
}
