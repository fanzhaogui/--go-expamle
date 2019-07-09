package main

import (
	"api.service.com/lib"
	"fmt"
	"strings"
)

func Test() {
	fmt.Println(strings.TrimSpace(" \t\n      a lone gopher         \n\t\r\n")) // a lone gopher

	originalText := "106"
	fmt.Println(originalText)
	mytext := []byte(originalText)

	cryptoText, _ := lib.Encrypt(mytext)
	fmt.Println(cryptoText)
	decryptedText, _ := lib.Decrypt(cryptoText)
	fmt.Println(decryptedText)
}
