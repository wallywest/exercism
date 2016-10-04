package raindrops

import "strconv"

const testVersion = 2

func Convert(num int) string {
	output := ""

	if num%3 == 0 {
		output += "Pling"
	}

	if num%5 == 0 {
		output += "Plang"
	}

	if num%7 == 0 {
		output += "Plong"
	}

	if output == "" {
		return strconv.Itoa(num)
	}

	return output
}

// The test program has a benchmark too.  How fast does your Convert convert?
