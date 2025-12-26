// Use this command to parse files searching for parameter match and print it to stdout
// Default parameter: CRITICAL
// Sample trigger commands:
//    go run .
//    go run . -level DEBUG

package main

import (
	"bufio"   // Provides buffered I/O, useful for efficient reading of files line by line
	"flag"    // Implements command-line flag parsing (e.g., -level DEBUG)
	"fmt"     // Implements formatted I/O (for printing output)
	"log"     // Implements simple logging, used here for error handling
	"os"      // Provides a platform-independent interface to operating system functionality (like file access)
	"strings" // Implements simple functions for manipulating strings (like checking for substrings)
)

func main() {
	// 1. Command-Line Flag Setup
	// flag.String defines a string flag named "level".
	// The second argument ("CRITICAL") is the default value if the flag is not provided.
	// The third argument is the usage message printed if the user asks for help.
	level := flag.String("level", "CRITICAL", "Log level to filter for")

	// flag.Parse executes the command-line parsing.
	// It must be called before the flag variables are accessed.
	flag.Parse()
	// 2. File Opening
	// os.Open attempts to open the file named "log.txt".
	file, err := os.Open("./logs.txt")

	// Standard error checking: if an error occurred (e.g., file not found),
	// log.Fatal prints the error and exits the program immediately.
	if err != nil {
		log.Fatal(err)
	}

	// defer schedules a function call (file.Close()) to be run after
	// the surrounding function (main) finishes executing. This ensures the file
	// is always closed, even if a runtime error occurs.
	defer file.Close()

	// 3. Buffered Reader Initialization
	// bufio.NewReader wraps the file handle, providing a buffered reading mechanism.
	// This makes reading line-by-line much more efficient than reading directly from the OS.
	bufReader := bufio.NewReader(file)

	// 4. File Iteration Loop
	// This is a classic Go idiom for reading a file stream until the end.

	// Initialization (before the first check): Read the first line of the file.
	// '\n' specifies the delimiter; ReadString reads until the next newline character.
	// 'line' holds the content, 'err' holds the error (e.g., io.EOF at end of file).
	for line, err := bufReader.ReadString('\n');

	// Condition (checked before each loop): Continue looping as long as no error has occurred.
	// When the end of the file is reached, err will be io.EOF, and the loop will stop.
	err == nil;

	// Post-iteration (run at the end of the block): Read the next line for the next iteration.
	line, err = bufReader.ReadString('\n') {

		// 5. Filtering and Output
		// The * operator dereferences the 'level' pointer to get its actual string value.
		// strings.Contains checks if the current 'line' contains the filter log level (e.g., "CRITICAL").
		if strings.Contains(line, *level) {

			// fmt.Print writes the matching line to standard output (stdout).
			// Since ReadString('\n') includes the newline, we use Print instead of Println.
			fmt.Print(line)
		}
	}
}
