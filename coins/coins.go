package coins

import (
	"bufio"
	. "fmt"
	"os"
)

func Pay() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var typesOfCoins int
	var total int
	min := 1
	var valuesOfCoins []int
	Fscan(in, &typesOfCoins, &total)
	for i := 0; i < typesOfCoins; i++ {
		var multiplier int
		Fscan(in, &multiplier)
		valuesOfCoins = append(valuesOfCoins, multiplier*min)
		min = valuesOfCoins[i]
	}

	var numberOfCoins int
	sub := 1
	max := valuesOfCoins[len(valuesOfCoins)-sub]
	for total <= 0 {
		if total-max < 0 {
			sub++
			if sub < len(valuesOfCoins) {
				max = valuesOfCoins[len(valuesOfCoins)-sub]
			}
		}
		total = total - max

		Fprintln(out, total)
		numberOfCoins++
	}
	Fprintln(out, numberOfCoins)
}
