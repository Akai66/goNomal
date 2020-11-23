//map里面嵌套map

package main

import (
	"fmt"
)

var graph = make(map[string]map[string]bool)

func main() {
	fmt.Println(hasEdge("a", "b"))
	addEdge("a", "b")
	fmt.Println(hasEdge("a", "b"))
}

func addEdge(from, to string) {
	if edges, ok := graph[from]; !ok {
		edges = make(map[string]bool)
		graph[from] = edges
	}
	graph[from][to] = true
}

func hasEdge(from, to string) bool {
	return graph[from][to] //对nil进行map取值，不会报异常，可以直接判断
}
