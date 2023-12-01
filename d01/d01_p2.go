package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
    file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	var res int
	scanner := bufio.NewScanner(file)

	digits := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9}

	for scanner.Scan() {
		sS := len(scanner.Text())
		sSV := -1
		eS := -1
		eSV := -1
		var cur int

		for i, c := range scanner.Text() {
			for k, v := range digits {
                found_idx := strings.HasPrefix(scanner.Text()[i:], k)
				if found_idx {
					if i < sS {
						sS = i
						sSV = v * 10
					}
					if i > eS {
						eS = i
						eSV = v
					}
				}
			}

			k := len(scanner.Text()) - 1 - i
			if parsed, err := strconv.Atoi(string(c)); err == nil {
				if i < sS {
					sS = i
					sSV = parsed * 10
				}
			}

			if parsed, err := strconv.Atoi(string(scanner.Text()[k])); err == nil {
				if k > eS {
					eS = k
					eSV = parsed
				}
			}

			if !(sSV == -1 && eSV == -1) {
				if sSV == -1 {
					sSV = eSV * 10
					sS = eS
				} else if eSV == -1 {
					eSV = sSV / 10
					eS = sS
				}
			}
		}

		cur = sSV + eSV
		fmt.Println(cur)
		res += cur
	}

	fmt.Println(res)
}
