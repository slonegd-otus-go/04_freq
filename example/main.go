package main

import (
	"fmt"

	freq "github.com/slonegd-otus-go/04_freq"
)

func main() {
	out := freq.Calculate(
		`text with many repeat words. many many, but not mutch!
	text with tabs. is it for test?
	yes, it's for tests`,
		4,
	)
	fmt.Printf("%v\n", out)

}
