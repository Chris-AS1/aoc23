package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
)

type camelCard struct {
	play  string
	class int
	bet   int
}

var ranking []rune = []rune{'A', 'K', 'Q', 'J', 'T', '9', '8', '7', '6', '5', '4', '3', '2'}

func countArray(s string) []int {
	count := make(map[rune]int)
	for _, c := range s {
		count[c] += 1
	}

	vec := make([]int, len(s))
	for i, c := range s {
		vec[i] = count[c]
	}

	return vec
}

func classify(play string) int {
	count := countArray(play)
	if slices.Contains(count, 5) {
		return 6
	}
	if slices.Contains(count, 4) {
		return 5
	}
	if slices.Contains(count, 3) {
		if slices.Contains(count, 2) {
			return 4
		}
		return 3
	}

	twos := 0
	for _, x := range count {
		if x == 2 {
			twos += 1
		}
	}
	if twos == 4 {
		return 2
	}

	if slices.Contains(count, 2) {
		return 1
	}

	return 0
}

func parseLine(line string) camelCard {
	split := strings.Split(line, " ")
	play_str, bet_str := split[0], split[1]
	rank := classify(play_str)
	bet, _ := strconv.Atoi(bet_str)
	return camelCard{play: play_str, class: rank, bet: bet}
}

func main() {
	file, err := os.Open("input.txt")
	weights := []camelCard{}

	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		card := parseLine(line)
		weights = append(weights, card)
	}

	sort.Slice(weights, func(i, j int) bool {
		// first check the class
		if weights[i].class != weights[j].class {
			return weights[i].class < weights[j].class
		}

		// then sort on semi-alphabetical
		var w, z int
		l, r := 0, 0
		for weights[i].play[l] == weights[j].play[r] &&
			l < len(weights[i].play) && r < len(weights[j].play) {
			l += 1
			r += 1
		}
		w = len(ranking) - slices.Index(ranking, rune(weights[i].play[l]))
		z = len(ranking) - slices.Index(ranking, rune(weights[j].play[r]))
		return w < z
	})

	res := 0
	for i, card := range weights {
		res += (i + 1) * card.bet

	}
	fmt.Println(res)
}
