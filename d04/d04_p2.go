package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
func parseLine(line string) (c_id int, score int) {
	split := strings.Split(line, ":")
	card, numbers := split[0], split[1]

	fmt.Sscanf(card, "Card %d", &c_id)
	split = strings.Split(numbers, "|")
	winning, given := split[0], split[1]

	winning_set := map[int]bool{}
	for _, x := range strings.Split(winning, " ") {
		x = strings.TrimSpace(x)
		if x == "" {
			continue
		}
		y, _ := strconv.Atoi(x)
		winning_set[y] = true
	}

	n := 0
	for _, x := range strings.Split(given, " ") {
		x = strings.TrimSpace(x)
		if x == "" {
			continue
		}
		y, _ := strconv.Atoi(x)
		_, ok := winning_set[y]
		if ok {
			n += 1
		}
	}
	return c_id, n
}

func main() {
	file, err := os.Open("input.txt")
	res := 0

	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	copies := map[int]int{}

	for scanner.Scan() {
		line := scanner.Text()
		g_id, p := parseLine(line)
		copies[g_id] += 1
		for i := 1; i < p+1; i += 1 {
			if _, ok := copies[g_id+i]; !ok {
				copies[g_id+i] = 0
			}
			copies[g_id+i] += copies[g_id]
		}

		res += copies[g_id]
	}
	fmt.Println(res)
}
