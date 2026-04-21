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
)

func main() {

	//Call count function
	fmt.Println(count(os.Stdin))

}

func count(r io.Reader) int {

	// Scanner used to read text from a Reader
	scanner := bufio.NewScanner(r)

	//Using Split to split type to words (default by lines)
	scanner.Split(bufio.ScanWords)

	//Defining the counter
	wordcount := 0

	// Increment the counter for word scanned
	for scanner.Scan() {
		wordcount++
	}

	return wordcount

}
