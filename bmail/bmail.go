package bmail

import (
	"bufio"
	. "fmt"
	"os"
	"sort"
)

func RouterPath() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	var numberOfRouters int
	Fscanln(in, &numberOfRouters)

	var routerNet = make(map[int]int)
	for i := 2; i <= numberOfRouters; i++ {
		var routerParentIndex int
		Fscan(in, &routerParentIndex)
		routerNet[i] = routerParentIndex
		// routerNet = append(routerNet, fmt.Sprintf("%d -> %d\n", i, routerParentIndex))
	}

	var output []int
	current := numberOfRouters
	for current != 1 {
		output = append(output, current)
		next := routerNet[current]
		current = next
	}
	output = append(output, 1)
	_ = sort.Reverse(sort.IntSlice(output))
	for i := 0; i < len(output); i++ {
		index := len(output) - 1
		Fprintf(out, "%d ", output[index-i])
	}
}
