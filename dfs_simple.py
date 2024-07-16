

NODES = 7

g = [[] for _ in range(NODES)]
visited = [False for _ in range(NODES)]

def add_edge(a, b):
	g[a].append(b)

def DFS(node):
	visited[node] = True
	print(f"visiting node {node}")

	for viz in g[node]:
		if not visited[viz]:
			DFS(viz)


def main():
	N, M = map(int, input().split())

	for i in range(M):
		a, b = map(int, input().split())
		add_edge(a, b)

	for i in range(1, N + 1):
		if not visited[i]:
			print(f"\nstart visit at {i}\n")
			DFS(i)


if __name__ == '__main__':
	main()

