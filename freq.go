package freq

import (
	"sort"
	"strings"
)

func Calculate(in string, qty int) []string {
	separators := []string{",", ".", "?", "!", "\t", "\n"}
	for _, separator := range separators {
		in = strings.Replace(in, separator, " ", -1)
	}

	words := strings.Split(in, " ")

	wordsQty := map[string]int{}
	for _, word := range words {
		if word != "" {
			wordsQty[word]++
		}
	}

	elements := slice(wordsQty)
	sort.Slice(elements, func(i, j int) bool {
		return elements[i].qty > elements[j].qty
	})

	res := make([]string, 0, qty)
	for i, element := range elements {
		res = append(res, element.word)
		if i == qty-1 {
			break
		}
	}
	return res
}

type element struct {
	word string
	qty  int
}

func slice(in map[string]int) []element {
	res := make([]element, 0, len(in))
	for word, qty := range in {
		res = append(res, element{word, qty})
	}
	return res
}
