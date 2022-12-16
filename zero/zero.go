package zero

import (
	"bufio"
	. "fmt"
	"os"
)

func zerinho() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	var a, b, c int
	Fscan(in, &a, &b, &c)

	if a == b && b == c {
		Fprintln(out, "*")
	} else if a == b {
		Fprintln(out, "C")
	} else if b == c {
		Fprintln(out, "A")
	} else {
		Fprintln(out, "B")
	}
}

// O(1)
