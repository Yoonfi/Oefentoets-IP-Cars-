package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var running = true

func main() {
	for running {
		PrintMenu()
	}
}

func PrintMenu() {
	fmt.Println("----- Bibliotheek ----")
	fmt.Println("1) Toon alle boeken")
	fmt.Println("2) Toon leengeschiedenis van klant")
	fmt.Println("3) Toon beschikbare boeken")
	fmt.Println("4) Afsluiten")
}

func ReadLine() string {
	buffer := bufio.NewReader(os.Stdin)
	input, _ := buffer.ReadString('\n')
	input = strings.TrimSuffix(input, "\n")
	input = strings.TrimSuffix(input, "\r")
	return input
}

func ReadFileToBytes(filename string) ([]byte, error) {
	return os.ReadFile(filename)
}
