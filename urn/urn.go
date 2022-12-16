package urn

import (
	"bufio"
	. "fmt"
	"os"
)

func Guess() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var colors int

	Fscanln(in, &colors)

	max := 0
	for i := 1; i <= colors; i++ {
		var numberOfColor int
		Fscan(in, &numberOfColor)
		if numberOfColor > max {
			max = numberOfColor
		}
	}

	Fprintln(out, max)
}
