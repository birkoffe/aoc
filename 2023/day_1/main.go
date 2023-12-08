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

func solve_part2(scanner *bufio.Scanner) int {
	digits := map[string]string{
		"1":     "1",
		"2":     "2",
		"3":     "3",
		"4":     "4",
		"5":     "5",
		"6":     "6",
		"7":     "7",
		"8":     "8",
		"9":     "9",
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}

	var sum int

	for scanner.Scan() {
		var a, b string

		for idx := range scanner.Text() {
			for key, value := range digits {
				if len(key) <= len(scanner.Text()[idx:]) {
					if scanner.Text()[idx:idx+len(key)] == key {
						if a == "" {
							a = value
						}

						b = value
					}
				}
			}
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
	fmt.Println("2023d1")

	files := []string{"sample_part_1.txt", "input.txt"}
	for _, file := range files {
		f, err := os.Open(file)
		if err != nil {
			fmt.Printf("Fail to open input.txt: %s\n", err)
			os.Exit(1)
		}
		defer f.Close()

		scanner := bufio.NewScanner(f)
		result := solve_part1(scanner)

		fmt.Printf("%20s = %d\n", file, result)
	}

	files = []string{"sample_part_2.txt", "input.txt"}
	for _, file := range files {
		f, err := os.Open(file)
		if err != nil {
			fmt.Printf("Fail to open input.txt: %s\n", err)
			os.Exit(1)
		}
		defer f.Close()

		scanner := bufio.NewScanner(f)
		result := solve_part2(scanner)

		fmt.Printf("%20s = %d\n", file, result)
	}
}
