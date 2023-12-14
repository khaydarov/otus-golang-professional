package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

type Entrance struct {
	Word  string
	Count uint
}

func Top10(s string) []string {
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

	entrances := make([]*Entrance, 0, len(m))
	for _, entrance := range m {
		entrances = append(entrances, entrance)
	}

	sort.Slice(entrances, func(i int, j int) bool {
		return entrances[i].Count > entrances[j].Count ||
			(entrances[i].Count == entrances[j].Count && entrances[i].Word < entrances[j].Word)
	})

	var result []string
	for i := 0; i < 10; i++ {
		result = append(result, entrances[i].Word)
	}

	return result
}

func processWord(word string) string {
	word = strings.ToLower(word)
	word = strings.TrimSpace(word)
	word = strings.Trim(word, ".,'!")

	return word
}
