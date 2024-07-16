package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func printf(f string, a ...interface{}) { fmt.Fprintf(writer, f, a...) }
func scanf(f string, a ...interface{})  { fmt.Fscanf(reader, f, a...) }

/*
input - quantidade de vértices e arestas
vértices convencionados numerados de 1 a N

N, M
1 2
2 3
.
.
.
*/

func add_edge(g [][]int, a, b int) {
	g[a] = append(g[a], b)
}

func print_graph(g [][]int, N int) {
	for i := 1; i <= N; i++ {
		printf("node %d -- ", i)
		for _, node := range g[i] {
			printf("%d ", node)
		}
		printf("\n")
	}
}

func main() {
	defer writer.Flush()

	var N, M int
	scanf("%d %d\n", &N, &M)

	g := make([][]int, N + 1)

	for i := 0; i < M; i++ {
		var a, b int
		scanf("%d %d\n", &a, &b)
		add_edge(g, a, b)
		// add_edge(g, b, a) // descomentar se o grafo for não direcionado
	}

	print_graph(g, N)
}