package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// bufio.NewReader – reads the full sentence (including spaces)
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter a sentence:")
	sentence, _ := reader.ReadString('\n')

	// strings.ToLower – ensures case-insensitive counting.
	//strings.Fields – splits the sentence into words.
	words := strings.Fields(strings.ToLower(sentence))

	// map[string]int – stores each word as a key and its count as a value.
	wordCount := make(map[string]int)

	for _, word := range words {
		wordCount[word]++
	}
	fmt.Println("\nWord Frequency:")
	for word, count := range wordCount {
		fmt.Printf("%s: %d\n", word, count)
	}
}
