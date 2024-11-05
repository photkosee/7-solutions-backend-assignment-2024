package main

import (
	"fmt"
	"strconv"
)

func decode(encoded string) (string, error) {
	pwd := make([]int, len(encoded)+1)

	for i, char := range encoded {
		// =: Sets the next value in pwd to be equal to the current value.
		// L: Sets the next value to 0 if the current value is non-zero; otherwise, it sets it to -1.
		// R: Increases the next value by 1 compared to the current value.
		switch char {
		case '=':
			pwd[i+1] = pwd[i]
		case 'L':
			if pwd[i] > 0 {
				pwd[i+1] = 0
			} else {
				pwd[i+1] = -1
			}
		case 'R':
			pwd[i+1] = pwd[i] + 1
		default:
			return "", fmt.Errorf("invalid character '%c' in input: only 'R', 'L', and '=' are accepted", char)
		}

		// If the current value is negative, increase it by 1 and adjust the previous values
		if pwd[i+1] < 0 {
			pwd[i+1]++

			// Iterate backwards
			for j := i; j >= 0; j-- {
				if encoded[j] == 'L' && pwd[j] <= pwd[j+1] {
					pwd[j]++
				} else if encoded[j] == 'R' && pwd[j] >= pwd[j+1] {
					break
				} else if encoded[j] == '=' && pwd[j] != pwd[j+1] {
					pwd[j] = pwd[j+1]
				}
			}
		}
	}

	result := ""
	for _, char := range pwd {
		result += strconv.Itoa(char)
	}

	return result, nil
}

func main() {
	var encoded string
	fmt.Print("Enter the encoded string: ")
	fmt.Scanln(&encoded)

	decoded, err := decode(encoded)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Decoded number set with minimum sum:", decoded)
	}
}
