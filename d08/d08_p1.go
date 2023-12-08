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
	cur := "AAA"
	cur_mov_idx := 0
	for cur != "ZZZ" {
		lr := MoveFromRune(rune(mov[cur_mov_idx]))
		cur = adj[cur][lr]

		res += 1
		cur_mov_idx = (cur_mov_idx + 1) % len(mov)
	}
    /* for k, v := range adj {
        fmt.Println(k, v)
    } */
	fmt.Println(res)
}
