package freq

import (
	"sort"
	"strings"
)

func Calculate(in string, qty int) []string {
	words := strings.Split(in, " ")
	wordsQty := map[string]int{}
	for _, word := range words {
		wordsQty[word]++
	}
	nodes := slice(wordsQty)
	sort.Slice(nodes, func(i, j int) bool {
		return nodes[i].qty < nodes[i].qty
	})
	res := make([]string, 0, qty)
	for i, node := range nodes {
		res = append(res, node.word)
		if i == qty-1 {
			break
		}
	}
	return res
}

type node struct {
	word string
	qty  int
}

func slice(in map[string]int) []node {
	res := make([]node, 0, len(in))
	for word, qty := range in {
		res = append(res, node{word, qty})
	}
	return res
}
