package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"unicode"
)

func main() {
	file, err := os.Open("calibration-codes.txt")
	if err != nil {
		log.Fatal("Can't open file: ", err)
	}
	defer file.Close()

	var total int32
	total = 0
	s := bufio.NewScanner(file)
	if err := s.Err(); err != nil {
		log.Fatal("Can't read the file: ", err)
	}

	for s.Scan() {
		line := s.Text()
		var firstDigit, lastDigit rune
		for _, char := range line {
			if unicode.IsDigit(char) {
				if firstDigit == 0 {
					firstDigit = char
				}
				lastDigit = char
			}
		}

		if firstDigit == 0 || lastDigit == 0 {
			codeValue := int32(firstDigit - '0')
			total += codeValue
			fmt.Println(codeValue)
		} else {
			firstDigitInt := int32(firstDigit - '0')
			lastDigitInt := int32(lastDigit - '0')
			codeValue := firstDigitInt*10 + lastDigitInt
			total += codeValue
			fmt.Println(codeValue)
		}
	}
	fmt.Println(total)
}
