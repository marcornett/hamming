/*
	package hamming checks for the Hamming Distance,
	the amount of mistakes occurring between two strands of DNA.

	func Distance will take in two strings and return a number.
		- equal length of strings will return Hamming Distance
		- unequal length will return 0
*/

package hamming

import "errors"

func Distance(a, b string) (int, error) {
	var hammingDistance int = 0

	if len(a) != len(b) {
		return hammingDistance, errors.New("Improper length")
	}

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			hammingDistance++
		}
	}
	return hammingDistance, nil
}
