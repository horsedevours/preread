package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"

	bolt "go.etcd.io/bbolt"
)

func main() {
	fmt.Println("getting list of common words...")
	file, err := os.Open("assets/3000-most-common-english.txt")
	if err != nil {
		fmt.Println("error opening file: ", err)
		os.Exit(1)
	}

	scanner := bufio.NewScanner(file)
	words := make([][]byte, 3000)
	for i := 0; i < 3000; i++ {
		scanner.Scan()
		line := make([]byte, len(scanner.Bytes()))
		copy(line, scanner.Bytes())
		words[i] = line
	}

	fmt.Println("creating Bolt file...")
	db, err := bolt.Open("preread.db", 0600, nil)
	if err != nil {
		fmt.Println("error creating DB: ", err)
		os.Exit(1)
	}
	defer db.Close()

	fmt.Println("adding words to CommonWords bucket...")
	err = db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte("CommonWords"))
		if err != nil {
			return fmt.Errorf("error creating the CommonWords bucker: %w", err)
		}

		for _, word := range words {
			err = b.Put(word, []byte{})
			if err != nil {
				return errors.New("error inserting word, aborting DB setup")
			}
		}

		return nil
	})
	if err != nil {
		fmt.Println("error populating database: ", err)
		removeDB()
		os.Exit(1)
	}

	fmt.Println("validate all 3000 words added...")
	err = db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("CommonWords"))

		if b.Stats().KeyN != 3000 {
			return fmt.Errorf("failed to insert 3000 common words, aborting DB setup")
		}

		return nil
	})

	if err != nil {
		fmt.Println("DB failed validation: ", err)
		removeDB()
		os.Exit(1)
	}

	fmt.Println("setup completed successfully")
	os.Exit(0)
}

func removeDB() {
	err := os.Remove("preread.db")
	if err != nil {
		fmt.Println("error removing broken DB")
	}
	fmt.Println("setup failed, broken DB removed")
}
