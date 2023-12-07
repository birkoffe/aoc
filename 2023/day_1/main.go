package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solve_part1(scanner *bufio.Scanner) int {
	var sum int

	for scanner.Scan() {
		var a, b string

		for _, s := range scanner.Text() {
			_, err := strconv.Atoi(string(s))
			if err != nil {
				continue
			}

			if a == "" {
				a = string(s)
				continue
			}

			b = string(s)
		}

		if b == "" {
			b = a
		}

		x, err := strconv.Atoi(a + b)
		if err != nil {
			fmt.Printf("Error parsing %s: %s\n", a+b, err)
			os.Exit(1)
		}

		sum += x
	}

	return sum
}

func main() {
	files := []string{"sample_part_1.txt", "input_part_1.txt"}
	for _, file := range files {
		f, err := os.Open(file)
		if err != nil {
			fmt.Printf("Fail to open input.txt: %s\n", err)
			os.Exit(1)
		}
		defer f.Close()

		scanner := bufio.NewScanner(f)
		result := solve_part1(scanner)

		fmt.Println("=", result)
	}
}
