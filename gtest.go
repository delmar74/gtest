package main

import (
	"bufio"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func LinesInFile(fileName string) []string {
	f, _ := os.Open(fileName)
	scanner := bufio.NewScanner(f)
	result := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		result = append(result, line)
	}
	return result
}

func main() {

	splitter := "|>"
  linebreak := "||"

	args := flag.String("f", "test.txt", "test file")
	flag.Parse()
	file := *args

	lines := LinesInFile(file)
	count := len(lines)

	for true {
		rand.Seed(time.Now().UnixNano())
		random := rand.Intn(count)

		line := strings.Split(lines[random], splitter)

		question := strings.Replace(line[0], linebreak, "\n", -1)
		answer := strings.Replace(line[1], linebreak, "\n", -1)

		fmt.Print("\033[H\033[2J")
		fmt.Println(question)
		bufio.NewReader(os.Stdin).ReadBytes('\n')
		fmt.Println("----------")
		fmt.Println(answer)
		fmt.Println("----------")
		bufio.NewReader(os.Stdin).ReadBytes('\n')
	}
}
