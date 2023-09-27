package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	var wordlists []string

	// Define command-line flags for wordlists
	flag.Var((*wordlistSlice)(&wordlists), "w", "Wordlist files in the format 'path:placeholder'")

	inputFlag := flag.String("input", "", "Input string with placeholders")
	flag.Parse()

	if *inputFlag == "" {
		log.Fatal("Please provide an input string with placeholders")
	}

	// Read and load wordlists
	wordlistData := make(map[string][]string)

	for _, wordlist := range wordlists {
		parts := strings.Split(wordlist, ":")
		if len(parts) != 2 {
			log.Fatalf("Invalid wordlist format: %s", wordlist)
		}

		path := parts[0]
		placeholder := parts[1]

		content, err := ioutil.ReadFile(path)
		if err != nil {
			log.Fatalf("Error reading %s: %v", path, err)
		}
		words := strings.Fields(string(content))
		wordlistData[placeholder] = words
	}

	// Generate wordlist combinations
	combinations := generateCombinations(*inputFlag, wordlistData)
	for _, combo := range combinations {
		fmt.Println(combo)
	}
}

type wordlistSlice []string

func (w *wordlistSlice) String() string {
	return fmt.Sprintf("%v", *w)
}

func (w *wordlistSlice) Set(value string) error {
	*w = append(*w, value)
	return nil
}

func generateCombinations(input string, wordlists map[string][]string) []string {
	combinations := []string{input}

	for placeholder, words := range wordlists {
		newCombinations := []string{}
		for _, combo := range combinations {
			for _, word := range words {
				newCombo := strings.Replace(combo, placeholder, word, -1)
				newCombinations = append(newCombinations, newCombo)
			}
		}
		combinations = newCombinations
	}

	return combinations
}
