package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func parse(line string) int {
	tot := 0
	for _, hash := range strings.Split(line, ",") {
		res := 0
		for _, c := range hash {
			res += int(c)
			res *= 17
			res %= 256
		}
		tot += res
	}
	return tot
}

func main() {
	file, err := os.Open("input.txt")

	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	res := 0
	for scanner.Scan() {
		res = parse(scanner.Text())
	}
	fmt.Println(res)
}
