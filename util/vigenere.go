package util

var ALPHABET_MAP = map[string]int{"A": 0, "B": 1, "C": 2, "D": 3, "E": 4, "F": 5, "G": 6, "H": 7, "I": 8, "J": 9, "K": 10, "L": 11, "M": 12, "N": 13, "O": 14, "P": 15, "Q": 16, "R": 17, "S": 18, "T": 19, "U": 20, "V": 21, "W": 22, "X": 23, "Y": 24, "Z": 25}
var ALPHABET = []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}

/*
Vigenere Cipher

Utilizes a Vigenere square which defines multiple cipher alphabets.
The order of the cipher alphabets in which to use for each of the plaintext letters
is defined by a keyword. The index of the keyword letter in the alphabet denotes which cipher row to use.

This was a method of encryption developed during the renissance which is impregnable to frequency analysis,
and foremost used technique by cryptanalysts of that day in age.
*/
func VigenereSquare() [][]string {
	// Generate Vigenere Square
	vigenereSquare := [][]string{}
	for i := 0; i < 26; i++ {
		row := []string{}
		for j := i; j < 26; j++ {
			if j < 25 {
				row = append(row, ALPHABET[j+1])
			}
		}
		for k := 0; k <= i; k++ {
			row = append(row, ALPHABET[k])
		}
		//uncomment to view square format
		// fmt.Println(row)
		vigenereSquare = append(vigenereSquare, row)
	}

	return vigenereSquare
}
