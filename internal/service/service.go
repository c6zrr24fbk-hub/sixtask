package service

import (
	"errors"
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func AutoDetectAndConvert(input string) (string, error) {
	if len(input) == 0 {
		return "", errors.New("empty input")
	}

	input = strings.TrimSpace(input)

	if isMorseCode(input) {
		return morse.ToText(input), nil
	} else {
		return morse.ToMorse(input), nil
	}
}

func isMorseCode(input string) bool {
	trimmed := strings.TrimSpace(input)
	if len(trimmed) == 0 {
		return false
	}

	allowedChars := ".-/ "
	morseChars := 0
	totalChars := 0

	for _, char := range trimmed {
		totalChars++
		if strings.ContainsRune(allowedChars, char) {
			morseChars++
		} else {
			if char != '\n' && char != '\r' && char != '\t' {
				return false
			}
		}
	}

	if morseChars == totalChars {
		return true
	}

	for _, char := range trimmed {
		if !strings.ContainsRune(".-/ \n\r\t", char) {
			return false
		}
	}

	hasMorseSymbol := false
	for _, char := range trimmed {
		if char == '.' || char == '-' {
			hasMorseSymbol = true
			break
		}
	}

	return hasMorseSymbol
}
