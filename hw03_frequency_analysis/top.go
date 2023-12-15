package hw03frequencyanalysis

import (
	"strings"
)

type Entrance struct {
	Word  string
	Count uint
}

func Top10(s string) []string {
	return topK(s, 10)
}

func topK(s string, k int) []string {
	if s == "" {
		return []string{}
	}

	words := strings.Fields(s)

	m := make(map[string]*Entrance)
	for _, word := range words {
		if word == "-" {
			continue
		}

		word = processWord(word)
		if v, ok := m[word]; ok {
			v.Count++
		} else {
			m[word] = &Entrance{
				Word:  word,
				Count: 1,
			}
		}
	}

	h := ConstructHeap(len(m))

	for _, entrance := range m {
		h.Insert(entrance)
	}

	var result []string
	for k > 0 && h.Size() > 0 {
		result = append(result, h.Extract().Word)
		k--
	}

	return result
}

func processWord(word string) string {
	word = strings.ToLower(word)
	word = strings.TrimSpace(word)
	word = strings.Trim(word, ".,'!")

	return word
}
