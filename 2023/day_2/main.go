package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		fmt.Printf("Fail to open input.txt: %s\n", err)
		os.Exit(1)
	}
	defer f.Close()

	var sum int64
	var max_cubes = map[string]int64{
		"red": 12, "green": 13, "blue": 14,
	}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()

		parsed_line := strings.Split(line, ": ")
		game_id, _ := strconv.ParseInt(strings.Split(parsed_line[0], " ")[1], 10, 64)

		valid := true

		for _, turn := range strings.Split(parsed_line[1], "; ") {
			for _, move := range strings.Split(turn, ", ") {
				m := strings.Split(move, " ")
				count, _ := strconv.ParseInt(m[0], 10, 64)
				if max_cubes[m[1]] < count {
					valid = false
				}
			}
		}

		if valid {
			sum += game_id
		}
	}

	fmt.Println(sum)
}
