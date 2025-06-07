package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"unicode"
)

func split(r rune) bool {
	switch r {
	case '-', '’':
		return false
	}
	return unicode.IsSpace(r) || unicode.IsPunct(r) || unicode.IsNumber(r)
}

func parseWords(b []byte) []string {
	words := bytes.FieldsFunc(b, split)

	strWords := make([]string, len(words))
	for i, word := range words {
		normWord := bytes.ToLower(word)
		strWords[i] = string(normWord)
	}
	return strWords
}

func parseText(rd io.Reader) {
	bufrd := bufio.NewReader(rd)

	buffer := make([]byte, 4096)
	_, err := bufrd.Read(buffer)
	if err != nil && err != io.EOF {
		log.Fatalf("error reading text: %s\n", err)
	}

	words := parseWords(buffer)

	fmt.Printf("Here's the words: %v", words)
}
