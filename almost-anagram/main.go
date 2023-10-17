package main

import (
	"fmt"
	"strconv"
)

func main() {
	var numberTemp string
	var number int
	fmt.Scanln(&numberTemp)
	var input []string
	number, err := strconv.Atoi(numberTemp)
	if err != nil {
		panic(err)
	}
	for i := 0; i < number; i++ {
		var temp string
		fmt.Scanln(&temp)
		input = append(input, temp)
	}

	// letters[0] = r,a,c,e
	// letters[1] = b,e,a,t
	var letters map[int][]rune = make(map[int][]rune)
	for i, v := range input {
		for _, v2 := range v {
			letters[i] = append(letters[i], v2)
		}
	}

	for i := 0; i < number; i++ {
		var almostAnagram int
		var firstLetters = letters[i] // grass
		for j := 0; j < number; j++ {
			if i == j {
				continue
			}
			if len(input[i]) != len(input[j]) {
				continue
			}
			var secondLetters = letters[j]
			for _, v := range firstLetters { // b r a s s
				var indexOfSecondLetter = -1
				for i2, v2 := range secondLetters { // g r a s s
					if v == v2 {
						indexOfSecondLetter = i2 // 1
						break
					}
				}
				if indexOfSecondLetter != -1 {
					var temp []rune
					for i2, v2 := range secondLetters { // g r a s s
						if i2 == indexOfSecondLetter { // 1
							continue
						}
						temp = append(temp, v2) // g a s s
					}
					secondLetters = temp // g a s s
				}
			}

			if len(secondLetters) == 1 {
				almostAnagram++
			}
		}
		fmt.Printf("%s %d\n", input[i], almostAnagram)
	}
}
