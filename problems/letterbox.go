package problems

import (
	"errors"
	"fmt"
)

func MakeLetterBox(word1 string, word2 string) (letterbox []string, err error) {
	// I can type here and the computer will do nothing

	// check if we can make a letterbox

	// last character in word 1 must be equal to first character in word 2
	lengthWord1 := len(word1)

	fmt.Println("length of word 1:", lengthWord1)
	fmt.Println("length of word 2:", len(word2))

	if word1[len(word1)-1] != word2[0] {
		errorMessage := "Last character of word 1 must equal first character of word 2"
		fmt.Println(errorMessage)
		return []string{}, errors.New(errorMessage)
	}

	// get unique characters from both words and ensure they number exactly 12

	// make an empty list to store unique characters
	uniques := map[string]bool{}

	// get a character from word
	for _, value := range word1 {
		character := string(value)
		fmt.Println(character)
		_, found := uniques[character]
		if found {
			continue
		} else {
			uniques[character] = true
		}
	}
	fmt.Println("unique characters from word 1:", uniques)
	fmt.Println("count of unique characters from word 1:", len(uniques))

	for _, value := range word2 {
		character := string(value)
		fmt.Println(character)
		_, found := uniques[character]
		if found {
			continue
		} else {
			uniques[character] = true
		}
	}
	fmt.Println("unique characters from both words:", uniques)
	fmt.Println("count of unique characters from both words:", len(uniques))

	if len(uniques) != 12 {
		return []string{}, errors.New("Total sum of unique characters between both words must be 12")
	}


	// check if the character from word 1 is in the list of unique characters
	// if the character from word 1 is in the list of unique characters, ignore and continue
	// otherwise, we found a new character, add it to the list of unique characters
	// repeat for word 2
	// check length of unique characters, if it is not 12, return an error

	// try to make a letterbox

	return []string{}, nil

}
