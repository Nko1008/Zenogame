package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var scanner = bufio.NewScanner(os.Stdin)

func readLine(prompt string) string {
	fmt.Print(prompt)
	scanner.Scan()
	return strings.TrimSpace(scanner.Text())
}

func readChoice(prompt string) int {
	for {
		input := readLine(prompt)
		n, err := strconv.Atoi(input)
		if err == nil {
			return n
		}
		fmt.Println("Entrée invalide, veuillez saisir un nombre.")
	}
}
