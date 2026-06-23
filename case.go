package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

// Mini Task 2: Tokenize the Text
func tokenize(text string) []string {
	return strings.Fields(text)
}

// Mini Task 3: Implement (up)
func ToUp(str string) string {
	tokens := tokenize(str)
	for i := 0; i < len(tokens); i++ {
		if tokens[i] == "(up)" {
			if i > 0 {
				tokens[i-1] = strings.ToUpper(tokens[i-1])
			}
			tokens = append(tokens[:i], tokens[i+1:]...)
			i--

		}

	}
	return strings.Join(tokens, " ")
}

// Mini Task 4: Implement (low)
func ToLow(str string) string {
	tokens := tokenize(str)
	for i := 0; i < len(tokens); i++ {
		if tokens[i] == "(low)" {
			if i > 0 {
				tokens[i-1] = strings.ToLower(tokens[i-1])
			}
			tokens = append(tokens[:i], tokens[i+1:]...)
			i--
		}
	}
	return strings.Join(tokens, " ")
}

func capitalize(str string) string {
	tokens := tokenize(str)
	for i, words := range tokens {
		runes := []rune(strings.ToLower(words))
		if len(runes) > 0 {
			runes[0] = unicode.ToUpper(runes[0])
		}
		tokens[i] = string(runes)

	}

	return strings.Join(tokens, " ")
}

// Mini Task 5: Implement (cap)
func ToCap(str string) string {
	tokens := tokenize(str)
	for i := 0; i < len(tokens); i++ {
		if tokens[i] == "(cap)" {
			if i > 0 {
				tokens[i-1] = capitalize(tokens[i-1])
			}
			tokens = append(tokens[:i], tokens[i+1:]...)
		}
	}
	return strings.Join(tokens, " ")
}

// Mini Task 6: Implement (up, n)
func ToUpN(str string) string {
	tokens := tokenize(str)
	for i := 0; i < len(tokens); i++ {
		if tokens[i] == "(up," || tokens[i] == "(low," || tokens[i] == "(cap," && i+1 < len(tokens) {
			numString := strings.TrimSuffix(tokens[i+1], ")")
			num, err := strconv.Atoi(numString)
			if err == nil {
				start := i - num
				if start < 0 {
					start = num
				}
				for j := start; j < i; j++ {
					switch tokens[i] {
					case "(up,":
						tokens[j] = strings.ToUpper(tokens[j])
					case "(low,":
						tokens[j] = strings.ToLower(tokens[j])
					case "(cap,":
						tokens[j] = capitalize(tokens[j])
					}

				}
				tokens = append(tokens[:i], tokens[i+2:]...)
				i--
			}
		}
	}
	return strings.Join(tokens, " ")

}

//Mini Task 7: Convert Binary Numbers

func main() {
	up := "Hello world (up)"
	low := "Hello WORLD (low)"
	fmt.Println(ToUp(up))
	fmt.Println(ToLow(low))
	fmt.Println("-------------------")
	cap := "Hello Mr michael (cap)"
	fmt.Println(ToCap(cap))
	fmt.Println(capitalize("unripe mangoes are here today"))
	fmt.Println("-------------------")
	fmt.Printf("%q\n", tokenize("hello (up, 2)"))
	upN := "This is so (cap, 2) exciting man woman boy girl (up, 4) man woman BOY GIRL (low, 2)"
	fmt.Println(ToUpN(upN))
}
