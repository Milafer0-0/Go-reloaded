package main

import (
	"regexp"
	"strconv"
	"strings"
)

func Transform(data []byte) []byte {
	lines := strings.Split(string(data), "\n")
	var processedLines []string

	for _, line := range lines {
		if strings.TrimSpace(line) == "" {
			processedLines = append(processedLines, "")
			continue
		}

		// Отделяем знаки препинания от скобок
		reSeparatePunct := regexp.MustCompile(`\)([\.\,\!\?\;\:])`)
		line = reSeparatePunct.ReplaceAllString(line, ") $1 ")

		// ОТДЕЛЯЕМ СЛИПШИЕСЯ МОДИФИКАТОРЫ (например, 7D(hex) -> 7D (hex) или (hex)(hex) -> (hex) (hex))
		reSeparateMods := regexp.MustCompile(`([^\s\(]+)(\((hex|bin|up|low|cap)(,[^\)]+)?\))`)
		line = reSeparateMods.ReplaceAllString(line, "$1 $2")
		reSeparateConsecutive := regexp.MustCompile(`\)\(`)
		line = reSeparateConsecutive.ReplaceAllString(line, ") (")

		words := strings.Fields(line)
		result := []string{}

		for i := 0; i < len(words); i++ {
			word := words[i]
			if word == "(hex)" {
				result = Hex(result)
				continue
			}
			if word == "(bin)" {
				result = Bin(result)
				continue
			}
			if word == "(up)" {
				result = Up(result, 1)
				continue
			}
			if word == "(low)" {
				result = Low(result, 1)
				continue
			}
			if word == "(cap)" {
				result = Cap(result, 1)
				continue
			}
			if word == "(up," && i+1 < len(words) {
				n := words[i+1]
				n = strings.TrimSuffix(n, ")")
				num, err := strconv.Atoi(n)
				if err != nil {
					result = append(result, word)
					continue
				}
				result = Up(result, num)
				i++
				continue
			}
			if word == "(low," && i+1 < len(words) {
				n := words[i+1]
				n = strings.TrimSuffix(n, ")")
				num, err := strconv.Atoi(n)
				if err != nil {
					result = append(result, word)
					continue
				}
				result = Low(result, num)
				i++
				continue
			}
			if word == "(cap," && i+1 < len(words) {
				n := words[i+1]
				n = strings.TrimSuffix(n, ")")
				num, err := strconv.Atoi(n)
				if err != nil {
					result = append(result, word)
					continue
				}
				result = Cap(result, num)
				i++
				continue
			}
			result = append(result, word)
		}

		finalLine := strings.Join(result, " ")
		finalLine = FixArticles(finalLine)
		finalLine = Punctuation(finalLine)
		finalLine = FixQuotes(finalLine)
		processedLines = append(processedLines, finalLine)
	}

	return []byte(strings.Join(processedLines, "\n"))
}

func Hex(result []string) []string {
	if len(result) == 0 {
		return result
	}
	last := result[len(result)-1]
	num, err := strconv.ParseInt(last, 16, 64)
	if err != nil {
		return result
	}
	result[len(result)-1] = strconv.Itoa(int(num))
	return result
}

func Bin(result []string) []string {
	if len(result) == 0 {
		return result
	}
	last := result[len(result)-1]
	num, err := strconv.ParseInt(last, 2, 64)
	if err != nil {
		return result
	}
	result[len(result)-1] = strconv.Itoa(int(num))
	return result
}

func Up(result []string, n int) []string {
	if len(result) == 0 {
		return result
	}
	if n > len(result) {
		n = len(result)
	}
	for u := 0; u < n; u++ {
		index := len(result) - 1 - u
		result[index] = strings.ToUpper(result[index])
	}
	return result
}

func Low(result []string, n int) []string {
	if len(result) == 0 {
		return result
	}
	if n > len(result) {
		n = len(result)
	}
	for u := 0; u < n; u++ {
		index := len(result) - 1 - u
		result[index] = strings.ToLower(result[index])
	}
	return result
}

func Cap(result []string, n int) []string {
	if len(result) == 0 {
		return result
	}
	if n > len(result) {
		n = len(result)
	}
	for u := 0; u < n; u++ {
		index := len(result) - 1 - u
		last := result[index]
		if len(last) == 0 {
			continue
		}
		runes := []rune(strings.ToLower(last))
		runes[0] = []rune(strings.ToUpper(string(runes[0])))[0]
		result[index] = string(runes)
	}
	return result
}

func Punctuation(text string) string {
	reBefore := regexp.MustCompile(`\s+([\.\,\!\?\;\:])`)
	text = reBefore.ReplaceAllString(text, "$1")

	reAfter := regexp.MustCompile(`([\.\,\!\?\;\:]+)([^ \.\,\!\?\;\:])`)
	text = reAfter.ReplaceAllString(text, "$1 $2")

	reSpaces := regexp.MustCompile(` +`)
	text = reSpaces.ReplaceAllString(text, " ")

	return strings.TrimSpace(text)
}

func FixArticles(text string) string {
	words := strings.Fields(text)
	vowels := "aeiouhAEIOUH"
	for i := 0; i < len(words)-1; i++ {
		current := words[i]
		if current == "a" || current == "A" {
			nextWord := words[i+1]
			cleanNext := strings.TrimLeft(nextWord, `.,!?;: '"`)
			if len(cleanNext) == 0 {
				continue
			}

			shouldBeAn := strings.ContainsRune(vowels, rune(cleanNext[0]))

			if shouldBeAn {
				if current == "a" {
					words[i] = "an"
				} else if current == "A" {
					words[i] = "An"
				}
			}
		}
	}
	return strings.Join(words, " ")
}

func FixQuotes(text string) string {
	re := regexp.MustCompile(`'\s+([^']*?)\s+'`)
	text = re.ReplaceAllString(text, "'$1'")

	text = regexp.MustCompile(`\s*'\s*([^'\s]+)\s*'\s*`).ReplaceAllString(text, " '$1' ")

	reSpaces := regexp.MustCompile(` +`)
	return strings.TrimSpace(reSpaces.ReplaceAllString(text, " "))
}
