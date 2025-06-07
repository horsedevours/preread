package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

/*
1. Open text file
2. Read contents
3. Parse contents for unique words
4. Check words against common words
5. Determine list of vocabulary words
6. Show to user
*/
func main() {
	file, err := os.Open("test.txt")
	if err != nil {
		log.Fatalf("Failed to open file: %v\n", err)
	}

	contents, _ := io.ReadAll(file)
	fmt.Printf("%s", contents)
}
