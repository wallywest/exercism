package hamming

import (
	"fmt"
	"strings"
)

const testVersion = 4

func Distance(a, b string) (int, error) {
	a_slice := strings.Split(a, "")
	b_slice := strings.Split(b, "")

	if len(a_slice) != len(b_slice) {
		return -1, fmt.Errorf("invalid length")
	}

	if a == b {
		return 0, nil
	}

	count := 0

	for i, valA := range a_slice {
		valB := b_slice[i]
		if valB != valA {
			count += 1
		}
	}

	return count, nil
}
