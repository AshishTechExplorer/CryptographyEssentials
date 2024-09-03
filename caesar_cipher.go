package main

import (
	"fmt"
	"strings"
)

// Encrypt performs the Caesar cipher encryption on the given plaintext with the specified shift.
func Encrypt(plaintext string, shift int) string {
	var encrypted strings.Builder

	for _, char := range plaintext {
		if char >= 'a' && char <= 'z' {
			shifted := (char-'a'+rune(shift))%26 + 'a'
			encrypted.WriteRune(shifted)
		} else if char >= 'A' && char <= 'Z' {
			shifted := (char-'A'+rune(shift))%26 + 'A'
			encrypted.WriteRune(shifted)
		} else {
			encrypted.WriteRune(char)
		}
	}

	return encrypted.String()
}

// Decrypt performs the Caesar cipher decryption on the given ciphertext with the specified shift.
func Decrypt(ciphertext string, shift int) string {
	return Encrypt(ciphertext, -shift) // Decryption is just encryption with the negative shift
}

func main() {
	plaintext := "Hello, World!"
	shift := 3

	encrypted := Encrypt(plaintext, shift)
	decrypted := Decrypt(encrypted, shift)

	fmt.Printf("Plaintext: %s\n", plaintext)
	fmt.Printf("Encrypted: %s\n", encrypted)
	fmt.Printf("Decrypted: %s\n", decrypted)
}
