package freq_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	freq "github.com/slonegd-otus-go/04_freq"
)

func TestCalculate(t *testing.T) {

	tests := []struct {
		name string
		in   string
		qty  int
		want []string
	}{
		{
			name: "some text",
			in:   "some text",
			qty:  2,
			want: []string{
				"some",
				"text",
			},
		},
		{
			name: "some text some another text text",
			in:   "some text some another text text",
			qty:  2,
			want: []string{
				"some",
				"text",
			},
		},
		{
			name: "separators",
			in: `text with many repeat words. many many, but not mutch!
	text with tabs. is it for test?
	yes, it's for tests`,
			qty: 4,
			want: []string{
				"many",
				"text",
				"with",
				"for",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := freq.Calculate(tt.in, tt.qty)
			assert.ElementsMatch(t, tt.want, got)
		})
	}
}
