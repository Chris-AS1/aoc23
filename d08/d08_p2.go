package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Move int64

const (
	Left Move = iota
	Right
)

func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)
	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}
	return result
}

func MoveFromRune(c rune) Move {
	switch c {
	case 'L':
		return Left
	case 'R':
		return Right
	default:
		return -1
	}
}

func findSteps(adj map[string][]string, cur string, mov string, curMoveIdx int) int {
	l := 0
	for cur[len(cur)-1] != 'Z' {
		l += 1
		curMove := MoveFromRune(rune(mov[curMoveIdx]))
		curMoveIdx = (curMoveIdx + 1) % len(mov)
		cur = adj[cur][curMove]
	}
	return l
}

func parseLine(line string) (node string, left string, right string) {
	split := strings.Split(line, " = ")
	node = split[0]

	lr_raw := split[1]
	lr_raw = strings.Trim(lr_raw, ")")
	lr_raw = strings.Trim(lr_raw, "(")
	lr_split := strings.Split(lr_raw, ", ")
	left, right = lr_split[0], lr_split[1]

	return node, left, right
}

func main() {
	file, err := os.Open("input.txt")

	if err != nil {
		panic(err)
	}
	defer file.Close()

	adj := map[string][]string{}

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	mov := scanner.Text()
	scanner.Scan()

	for scanner.Scan() {
		line := scanner.Text()
		n, l, r := parseLine(line)
		adj[n] = []string{l, r}
	}

	start_positions := []string{}
	for k := range adj {
		if k[len(k)-1] == 'A' {
			start_positions = append(start_positions, k)
		}
	}
	res := 1
	visited := make(map[string]int)
	for _, k := range start_positions {
		for k := range adj {
			visited[k] = -1
		}
		cyc := findSteps(adj, k, mov, 0)
		res = LCM(res, cyc)
	}
	fmt.Println(res)
}
