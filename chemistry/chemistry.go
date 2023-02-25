package chemistry

import (
	"bufio"
	. "fmt"
	"os"
)

func Mix() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	max := uint64(1)
	var numberOfReagents, numberOfReactions int
	Fscanln(in, &numberOfReagents, &numberOfReactions)

	var combinations [][]int // lista de arestas - arvore de extensão/geradora minima - min spanning tree
	for i := 0; i < numberOfReactions; i++ {
		var c1, c2 int
		Fscanln(in, &c1, &c2)

		combinations = append(combinations, []int{c1, c2})
		combinations = append(combinations, []int{c2, c1})
	}

	for i := 2; i <= numberOfReagents; i++ {
		for _, c := range combinations {
			if c[0] == i && c[1] < i {
				max *= 2
				break
			}
		}
	}
	Fprintln(out, max)
}
