package main

import (
	"fmt"
	"strings"
)

const originalLetter = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func hashLetterFn(key int, letter string) string {
	runes := []rune(letter)
	lastLetterKey := string(runes[len(letter)-key : len(letter)])
	leftOversLetter := string(runes[0 : len(letter)-key])
	return fmt.Sprintf("%s%s", lastLetterKey, leftOversLetter)
}

func shiftLetter(letter string, shift int) string {
	pos := strings.Index(originalLetter, letter)
	if pos != -1 {
		letterPosition := (pos + shift + len(originalLetter)) % len(originalLetter)
		return string(originalLetter[letterPosition])
	}
	return letter
}

func applyTransformation(text string, key int, transformation func(string, int) string) string {
	var result string
	findOne := func(r rune) rune {
		result += transformation(string([]rune{r}), key)
		return r
	}
	strings.Map(findOne, text)
	return result
}

func encrypt(key int, plainText string) string {
	return applyTransformation(plainText, key, shiftLetter)
}

func decrypt(key int, encryptedText string) string {
	return applyTransformation(encryptedText, -key, shiftLetter)
}

func main() {
	plainText := "HELLOWORLD" // Enter string to encrypt/decrypt
	fmt.Println("Plain Text", plainText)

	key := 5

	encrypted := encrypt(key, plainText)
	fmt.Println("Encrypted", encrypted)

	decrypted := decrypt(key, encrypted)
	fmt.Println("Decrypted", decrypted)
}
