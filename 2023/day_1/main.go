package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		fmt.Printf("Fail to open input.txt: %s\n", err)
		os.Exit(1)
	}
	defer f.Close()

	var sum int64

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		var a, b string

		for _, s := range scanner.Text() {
			_, err := strconv.ParseInt(string(s), 10, 64)
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

		x, err := strconv.ParseInt(a+b, 10, 32)
		if err != nil {
			fmt.Printf("Error parsing %s: %s\n", a+b, err)
			os.Exit(1)
		}

		sum += x
	}

	fmt.Println(sum)
}
