package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
func parseLine(line string) (game_id int, r int, g int, b int) {
	game, colors, found := strings.Cut(line, ":")
	if !found {
		return 0, 0, 0, 0
	}

	// game_id
	fmt.Sscanf(game, "Game %d", &game_id)

	// colors
	subsets := strings.Split(colors, ";")
	r, g, b = 0, 0, 0
	for _, set := range subsets {
		colors_to_parse := strings.Split(set, ", ")
		tr, tg, tb := 0, 0, 0
		for _, unk_color := range colors_to_parse {
			read := 0
			_, err := fmt.Sscanf(unk_color, "%d red", &read)
			if err == nil {
				tr += read
			}
			_, err = fmt.Sscanf(unk_color, "%d green", &read)
			if err == nil {
				tg += read
			}
			_, err = fmt.Sscanf(unk_color, "%d blue", &read)
			if err == nil {
				tb += read
			}
		}
		r = max(r, tr)
		g = max(g, tg)
		b = max(b, tb)
	}

	return game_id, r, g, b
}

func main() {
	file, err := os.Open("input.txt")
	res := 0

	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		game_id, r, g, b := parseLine(line)
		if r > 12 || g > 13 || b > 14 {
			continue
		}
		res += game_id
	}
	fmt.Println(res)
}
