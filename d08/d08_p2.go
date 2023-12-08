package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type camelCard struct {
	play  string
	class int
	bet   int
}

type Move int64

const (
	Left Move = iota
	Right
)

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

func allValid(pts []string) bool {
	for _, v := range pts {
		if v[len(v)-1] != 'Z' {
			return false
		}
	}
	return true
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

	res := 0
	cur := []string{}
	for k := range adj {
		if k[len(k)-1] == 'A' {
			cur = append(cur, k)
		}
	}

	cur_mov_idx := 0
	for !allValid(cur) {
		lr := MoveFromRune(rune(mov[cur_mov_idx]))
		for i, v := range cur {
			cur[i] = adj[v][lr]
		}
		res += 1
		cur_mov_idx = (cur_mov_idx + 1) % len(mov)
	}
	fmt.Println(res)
}
