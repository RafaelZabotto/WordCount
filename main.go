package main

/*
DISCLAIMER:
This program basicaly counts the words may received by a echo.
Like: echo "My first command line tool with Go" | ./wordCounter

The prime function is described above, therefore, everything beyond that is an update.
*/

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"flag"
)

func main() {

	//This mofication counts the numbers of lines instead words
	lines := flag.Bool("l", false, "Count lines")

	//Parsing the flags provided by the user
	flag.Parse()

	fmt.Println(count(os.Stdin, *lines))

}

func count(r io.Reader, countLines bool) int {

	// Scanner used to read text from a Reader
	scanner := bufio.NewScanner(r)

	// If the count lines flag is not set, we want to count words so we define
	// the scanner split type to words (default is split by lines)
	if !countLines {
		scanner.Split(bufio.ScanWords)
	}

	//Defining the counter
	wordcount := 0

	// Increment the counter for every word or line scanned
	for scanner.Scan() {
		wordcount++
	}

	return wordcount
}
