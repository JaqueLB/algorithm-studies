package numislands

const TERRA string = "1"
const AGUA string = "0"

// m linhas e n colunas
func numIslands(grid [][]byte) int {
	result := 0
	visited := createVisited(grid)
	for m := 0; m < len(grid); m++ {
		for n := 0; n < len(grid[m]); n++ {
			if string(grid[m][n]) == TERRA && !visited[m][n] {
				DFS(grid, visited, m, n)
				result++
			}
		}
	}
	return result
}

func createVisited(grid [][]byte) [][]bool {
	visited := make([][]bool, len(grid))
	for i := 0; i < len(grid); i++ {
		visited[i] = make([]bool, len(grid[i]))
	}
	return visited
}

var vizm = []int{1, 0, -1, 0}
var vizn = []int{0, 1, 0, -1}

func isIn(j int, k int, maxRow int, maxCol int) bool {
	return 0 <= j && j < maxRow && 0 <= k && k < maxCol
}

func DFS(grid [][]byte, visited [][]bool, m int, n int) {
	visited[m][n] = true

	for i := 0; i < len(vizm); i++ {
		j := m + vizm[i]
		k := n + vizn[i]
		if isIn(j, k, len(grid), len(grid[m])) && string(grid[j][k]) == TERRA && !visited[j][k] {
			DFS(grid, visited, j, k)
		}
	}
}

// O(mn) -> complexidade de tempo
// O(mn) -> complexidade de espaço
