package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	"os"

	"github.com/joho/godotenv"
	"golang.org/x/term"
)

func encrypt(password, key []byte) (string, error) {
	if len(key) != 32 {
		return "", fmt.Errorf("Key must be 32 bytes long")
	}

	init_vect := make([]byte, aes.BlockSize)
	if _, err := io.ReadFull(rand.Reader, init_vect); err != nil {
		return "", err
	}

	padding := aes.BlockSize - len(password)%aes.BlockSize

	password = append(password, make([]byte, padding)...)
	for i := range password[len(password)-padding:] {
		password[len(password)-padding+i] = byte(padding)
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	ciphertext := make([]byte, len(password))
	mode := cipher.NewCBCEncrypter(block, init_vect)
	mode.CryptBlocks(ciphertext, password)

	return hex.EncodeToString(append(init_vect, ciphertext...)), nil
}

func main() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

    fmt.Println("Enter Your Password: ")
	bytePassword, err := term.ReadPassword(int(os.Stdin.Fd()))
	if err != nil {
		panic(err)
	}

	pass := string(bytePassword)

	keyHex := os.Getenv("KEY")

	_, found := os.LookupEnv("KEY")
	fmt.Printf("Key found: %t\n", found)

	key, err := hex.DecodeString(keyHex)
	if err != nil {
		panic("Invalid hex key")
	}

	encrypted, err := encrypt([]byte(pass), key)
	if err != nil {
		fmt.Printf("Error Encrypting: %v\n", err)
		return
	}

	fmt.Printf("Encrypted (hex): %s\n", encrypted)
}
