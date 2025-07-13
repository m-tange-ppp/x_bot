package calc

import (
	"fmt"
	"strconv"
	"testing"
)

func TestPrimeFactorize(t *testing.T) {
	tests := []struct {
		input    int
		expected []int
	}{
		{0, []int{}},
		{1, []int{}},
		{2, []int{2}},
		{12, []int{2, 2, 3}},
		{25, []int{5, 5}},
		{360, []int{2, 2, 2, 3, 3, 5}},
	}

	for _, tt := range tests {
		t.Run(strconv.Itoa(tt.input), func(t *testing.T) {
			factors := PrimeFactorize(tt.input)
			fmt.Printf("PrimeFactorize(%d) = %v", tt.input, factors)

			// 結果を検証
			if len(factors) != len(tt.expected) {
				t.Errorf("Expected %v, got %v", tt.expected, factors)
				return
			}

			for i, factor := range factors {
				if factor != tt.expected[i] {
					t.Errorf("Expected %v, got %v", tt.expected, factors)
					return
				}
			}
		})
	}
}
