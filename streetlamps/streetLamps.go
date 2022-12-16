package streetlamps

import (
	"bufio"
	. "fmt"
	"os"
)

func isLighted(street []rune, i int) bool {
	return street[i-1] == '*' || street[i] == '*' || street[i+1] == '*'
}

func MakeLight() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var tests int
	Fscanln(in, &tests)

	for i := 0; i < tests; i++ {
		var blocks int
		var street string
		Fscan(in, &blocks, &street)
		var lamps int

		streetRunes := []rune("." + street + ".")

		for j := 1; j <= blocks; j++ {
			if !isLighted(streetRunes, j) {
				streetRunes[j+1] = '*'
				lamps++
			}
		}

		Fprintln(out, lamps)
	}
}
