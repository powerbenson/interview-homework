package util

import (
	"fmt"
	"unicode"
	"strings"
	"strconv"
	"regexp"
)

func ValidateInput(str string) bool {
	regex := `^(-?\d+)(\s-?\d+)*$`
	match, _ := regexp.MatchString(regex, str)
	return match
}

func ParseIntArray(str string) ([]int, error) {
	strArr := strings.Split(str, " ")
	var intArr []int
	for _, s := range strArr {
		num, err := strconv.Atoi(s)
		if err != nil {
			return nil, fmt.Errorf("Invalid integer: %s", s)
		}
		intArr = append(intArr, num)
	}
	return intArr, nil
}

func TrimTrailingWhitespace(data []byte) string {
	trimmedData := strings.TrimRightFunc(string(data), unicode.IsSpace)
	return trimmedData
}