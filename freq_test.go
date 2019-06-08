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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := freq.Calculate(tt.in, tt.qty)
			assert.ElementsMatch(t, tt.want, got)
		})
	}
}
