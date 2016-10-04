package secret

import "sort"

func Handshake(input int) []string {

	result := make([]string, 0)

	if input <= 0 {
		return nil
	}

	bins := map[int]string{
		1: "wink",
		2: "double blink",
		4: "close your eyes",
		8: "jump",
	}

	keys := []int{1, 2, 4, 8}
	sort.Ints(keys)

	for _, val := range keys {
		bin := bins[val]

		if input&val == val {
			result = append(result, bin)
		}
	}

	if input&16 == 16 {
		return reverse(result)
	}

	return result
}

func reverse(a []string) []string {
	for i := len(a)/2 - 1; i >= 0; i-- {
		opp := len(a) - 1 - i
		a[i], a[opp] = a[opp], a[i]
	}

	return a
}
