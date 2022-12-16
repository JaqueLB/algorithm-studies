package dive

import (
	"bufio"
	. "fmt"
	"os"
)

func Dive() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	var sent, returned int
	var lista []int
	Fscan(in, &sent, &returned)

	for i := 0; i < returned; i++ {
		var value int
		Fscan(in, &value)
		lista = append(lista, value)
	}

	if len(lista) == sent {
		Fprintln(out, "*")
	} else {
		var divers []int
		for j := 1; j <= sent; j++ {
			divers = append(divers, j)
		}

		// check divers
		for _, code := range lista {
			index := findDiver(divers, code)
			if index != -1 {
				divers = append(divers[:index], divers[index+1:]...)
			}
		}

		for _, d := range divers {
			Fprintf(out, "%d ", d)
		}
		Fprintln(out, "")
	}

}

func findDiver(divers []int, diver int) int {
	for i, d := range divers {
		if d == diver {
			return i
		}
	}
	return -1
}
