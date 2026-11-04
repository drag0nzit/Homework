package main

import (
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
	"unicode"
	"unicode/utf8"
)

type WordCount struct {
	Word  string
	Count int
}

func GetWords(text string) []string {
	var words []string
	var current strings.Builder

	for _, r := range strings.ToLower(text) {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			current.WriteRune(r)
		} else if current.Len() > 0 {
			words = append(words, current.String())
			current.Reset()
		}
	}

	if current.Len() > 0 {
		words = append(words, current.String())
	}

	return words
}

func CountSentences(text string) int {
	count := 0
	inSentenceEnd := false

	for _, r := range text {
		if r == '.' || r == '!' || r == '?' {
			if !inSentenceEnd {
				count++
			}
			inSentenceEnd = true
		} else if !unicode.IsSpace(r) &&
			r != '"' && r != '\'' &&
			r != ')' && r != ']' {
			inSentenceEnd = false
		}
	}

	return count
}

func Analyze(text string) {
	words := GetWords(text)

	characters := len(text)
	runes := utf8.RuneCountInString(text)
	sentences := CountSentences(text)

	frequency := make(map[string]int)

	for _, word := range words {
		frequency[word]++
	}

	fmt.Println("\nКоличество слов:", len(words))
	fmt.Println("Количество символов:", characters)
	fmt.Println("Количество рун:", runes)
	fmt.Println("Количество предложений:", sentences)
	fmt.Println("Уникальных слов:", len(frequency))

	if len(words) == 0 {
		fmt.Println("Текст не содержит слов")
		return
	}

	longest := words[0]
	shortest := words[0]

	for _, word := range words {
		if utf8.RuneCountInString(word) >
			utf8.RuneCountInString(longest) {
			longest = word
		}

		if utf8.RuneCountInString(word) <
			utf8.RuneCountInString(shortest) {
			shortest = word
		}
	}

	fmt.Println("Самое длинное слово:", longest)
	fmt.Println("Самое короткое слово:", shortest)

	var popular string
	maxCount := 0

	for word, count := range frequency {
		if count > maxCount {
			maxCount = count
			popular = word
		}
	}

	fmt.Println("Самое частое слово:", popular)
	fmt.Println("Количество повторений:", maxCount)

	var result []WordCount

	for word, count := range frequency {
		result = append(result, WordCount{word, count})
	}

	sort.Slice(result, func(i, j int) bool {
		if result[i].Count == result[j].Count {
			return result[i].Word < result[j].Word
		}
		return result[i].Count > result[j].Count
	})

	fmt.Println("\nТоп-10 частых слов:")

	limit := 10
	if len(result) < limit {
		limit = len(result)
	}

	for i := 0; i < limit; i++ {
		fmt.Println(result[i].Word, "-", result[i].Count)
	}
}

func main() {
	fmt.Println("Введите текст. Для завершения нажмите Ctrl+D (Linux/macOS) или Ctrl+Z, затем Enter (Windows):")

	data, err := io.ReadAll(os.Stdin)
	if err != nil {
		fmt.Println("Ошибка чтения текста:", err)
		return
	}

	Analyze(string(data))
}
