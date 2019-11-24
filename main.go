package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func LinesInFile(fileName string) []string {
	f, _ := os.Open(fileName)
	// Create new Scanner.
	scanner := bufio.NewScanner(f)
	result := []string{}
	// Use Scan.
	for scanner.Scan() {
		line := scanner.Text()
		// Append line to result.
		result = append(result, line)
	}
	return result
}

func main() {

	//  file := "test.txt"
	file := "rhce.txt"
	splitter := "|>"

	lines := LinesInFile(file)
	count := len(lines)

	for true {
		rand.Seed(time.Now().UnixNano())
		random := rand.Intn(count)

		line := strings.Split(lines[random], splitter)
		question := strings.Replace(line[0], "||", "\n", -1)
		answer := strings.Replace(line[1], "||", "\n", -1)
		//fmt.Println("")
    fmt.Print("\033[H\033[2J")
		fmt.Println(question)
		bufio.NewReader(os.Stdin).ReadBytes('\n')
		fmt.Println("----------")
		fmt.Println(answer)
		fmt.Println("----------")
		bufio.NewReader(os.Stdin).ReadBytes('\n')
	}
}
