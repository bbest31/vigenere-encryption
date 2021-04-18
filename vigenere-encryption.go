package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/vigenere-cipher/v2/util"
)

func main() {

	fmt.Print("Enter phrase to encode: ")
	reader := bufio.NewReader(os.Stdin)
	// ReadString will block until the delimiter is entered
	plaintext, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("An error occured while reading input. Please try again", err)
		return
	}

	// remove the delimeter from the string
	plaintext = strings.TrimSuffix(plaintext, "\n")
	plaintext = strings.ToUpper(plaintext)

	fmt.Print("Enter keyword: ")
	keyword, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("An error occured while reading the keyword. Please try again", err)
		return
	}

	keyword = strings.TrimSuffix(keyword, "\n")
	keyword = strings.ToUpper(keyword)

	encryptedMsg := encode(plaintext, keyword)
	fmt.Printf("Encrypted Message: %v\n", encryptedMsg)

	decryptedMsg := decode(encryptedMsg, keyword)
	fmt.Printf("Decrypted Message: %v\n", decryptedMsg)

}

func encode(plaintext string, keyword string) string {
	ciphertext := ""

	vigenereSquare := util.VigenereSquare()
	keyQueue := strings.Split(keyword, "")

	for _, letter := range plaintext {
		if string(letter) != " " && letter != 13 {
			// Get the cipher alphabet of the Vigenere Square to use based off the keyword character
			var cipherRow int = util.ALPHABET_MAP[keyQueue[0]]
			// Get the encrypted value
			encryptedValue := vigenereSquare[cipherRow][util.ALPHABET_MAP[string(letter)]]
			ciphertext += encryptedValue

			// Update keyQueue
			keyQueue = append(keyQueue[1:], keyQueue[0])

		} else if string(letter) == " " {
			ciphertext += " "
		}
	}

	return ciphertext
}

func decode(ciphertext string, keyword string) string {
	plaintext := ""

	vigenereSquare := util.VigenereSquare()
	keyQueue := strings.Split(keyword, "")

	for _, letter := range ciphertext {
		if string(letter) != " " {
			var cipherRow int = util.ALPHABET_MAP[keyQueue[0]]

			for i, v := range vigenereSquare[cipherRow] {
				if v == string(letter) {
					plaintext += util.ALPHABET[i]
					break
				}
			}

			// Update keyQueue
			keyQueue = append(keyQueue[1:], keyQueue[0])

		} else {
			plaintext += " "
		}
	}

	return plaintext
}
