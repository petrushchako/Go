// Use this command to parse files searching for parameter match and print it to stdout
// Default parameter: CRITICAL
// Sample trigger commands:
//    go run .
//    go run . -level DEBUG

package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	level := flag.String("level", "CRITICAL", "Log level to filter for")
	flag.Parse()

	file, err := os.Open("./log.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	bufReader := bufio.NewReader(file)

	for line, err := bufReader.ReadString('\n'); err == nil; line, err = bufReader.ReadString('\n') {
		if strings.Contains(line, *level) {
			fmt.Print(line)
		}
	}
}
