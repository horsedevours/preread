package main

import (
	"fmt"
	"log"
	"os"

	"go.etcd.io/bbolt"
)

func main() {
	db, err := bbolt.Open("setup/preread.db", 0600, nil)
	if err != nil {
		log.Fatalf("error opening DB: %v", err)
	}
	defer db.Close()

	// Open file
	file, err := os.Open("test.txt")
	if err != nil {
		log.Fatalf("Failed to open file: %v\n", err)
	}

	// Parse words from file
	words := parseText(file)

	// Keep only unique words
	uniqueWords := map[string]struct{}{}
	for _, word := range words {
		fmt.Printf("Next word: %q\n", word)
		uniqueWords[word] = struct{}{}
	}

	fmt.Println("Unique words: ", uniqueWords)
	fmt.Printf("Number unique words: %d\n", len(uniqueWords))

	// Get rid of known common words
	err = db.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte("CommonWords"))

		for w := range uniqueWords {
			if v := b.Get([]byte(w)); v != nil {
				delete(uniqueWords, w)
			}
		}

		return nil
	})
	if err != nil {
		log.Fatalln("error reading values from DB: ", err)
	}

	fmt.Println("Unique words: ", uniqueWords)

	// Make API calls to check rarity of words that are left

	// If word is common, add to list of common words

	// If word is not common, store with relevant information

	// Display vocabulary words to user
	// If lots, display a sample of the most difficult words

	// Let the user have a copy of all the words

	// After a text has been analyzed, let a user view those words
	// without having to re-analyze the text
}
