package prince

import (
	"bufio"
	"fmt"
	. "fmt"
	"os"
)

const LAND string = "."
const WATER string = "#"
const PRINCE string = "@"
const GRIDLIM int = 20

func isIn(h int, w int, maxHeight int, maxWidth int) bool {
	return 0 <= h && h < maxHeight && 0 <= w && w < maxWidth
}

func dfs(land [][]string, visited [][]bool, h int, w int) {
	var vizm = []int{1, 0, -1, 0}
	var vizn = []int{0, 1, 0, -1}
	visited[h][w] = true

	for i := 0; i < len(vizm); i++ {
		newHeight := h + vizn[i]
		newWidth := w + vizm[i]
		if isIn(newHeight, newWidth, len(land), len(land[h])) &&
			string(land[newHeight][newWidth]) == LAND &&
			!visited[newHeight][newWidth] {
			dfs(land, visited, newHeight, newWidth)
		}
	}
}

func VisitLand() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var tests int
	Fscanln(in, &tests)

	for t := 0; t < tests; t++ {
		var width, height int
		Fscanln(in, &width, &height)

		// create grid for land
		land := make([][]string, height)
		for h := 0; h < height; h++ {
			land[h] = make([]string, width)
		}

		// populate land
		for h := 0; h < height; h++ {
			var line string
			Fscanf(in, "%s\n", &line)
			for w := 0; w < width; w++ {
				land[h][w] = fmt.Sprintf("%c", line[w])
			}
		}

		// create grid for the patches already visited
		visited := make([][]bool, height)
		for h := 0; h < height; h++ {
			visited[h] = make([]bool, width)
		}

		// find the prince and start searching for patches
		for h := 0; h < height; h++ {
			for w := 0; w < width; w++ {
				if string(land[h][w]) == PRINCE {
					dfs(land, visited, h, w)
				}
			}
		}

		// count how many patches of land were visited
		result := 0
		for _, v := range visited {
			for _, l := range v {
				if l {
					result++
				}
			}
		}

		Fprintf(out, "Case %d: %d\n", t+1, result)
	}
}
