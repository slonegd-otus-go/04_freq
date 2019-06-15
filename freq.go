package freq

import (
	"sort"
	"strings"
)

func Calculate(in string, qty int) []string {
	words := strings.Fields(in)

	wordsQty := map[string]int{}
	for _, word := range words {
		wordsQty[word]++
	}

	elements := sliceElements(wordsQty)
	sort.Slice(elements, func(i, j int) bool {
		return elements[i].qty > elements[j].qty
	})

	return sliceStrings(elements, qty)
}

type element struct {
	word string
	qty  int
}

func sliceElements(in map[string]int) []element {
	res := make([]element, 0, len(in))
	for word, qty := range in {
		res = append(res, element{word, qty})
	}
	return res
}

func sliceStrings(elements []element, maxLen int) []string {
	res := make([]string, 0, maxLen)
	for i, element := range elements {
		res = append(res, element.word)
		if i == maxLen-1 {
			break
		}
	}
	return res
}
