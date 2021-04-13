package main

import (
	"fmt"
	"regexp"
	"strings"
)

func clean(word string) string {
	regex := regexp.MustCompile(`[^a-zA-Z]`)
	return regex.ReplaceAllString(word, "")
}

func reverse(word *string) string {
	splitWord := strings.Split(*word , "")
	temp := []string{}

	for i := len(splitWord)-1; i >= 0; i-- {
		temp = append(temp, splitWord[i]) 
	}

	reversedWord := strings.Join(temp, "")
	return reversedWord
}

func main() {
	input := []string{
		"Yo de todo te doy.",
		"A ti no bonita",
		"Esto no es un palindromo",
		"",
		"******* :::::",
	}

	for i, _ := range input{
		wordClened := clean(input[i])
		if len(wordClened) > 0{
			lowerString := strings.ToLower(wordClened)
			reversedWord:= reverse(&lowerString)
			result := reversedWord == lowerString

			if result {
				fmt.Printf("%t\t-> '%s' is a palindorme. \n", result, input[i])
			}else{
				fmt.Printf("%t\t-> '%s' is not a palindorme.\n", result, input[i])
			}
		}
	}
}
