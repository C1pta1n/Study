package main

import (
	"container/list"
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	executionTime := make([]int64, n)
	graph := make([][]int, n)
	inDegree := make([]int, n)

	for i := 0; i < n; i++ {
		var t int64
		fmt.Scan(&t)
		executionTime[i] = t

		for {
			var dep int
			if _, err := fmt.Scan(&dep); err != nil {
				break
			}
			dep--
			graph[dep] = append(graph[dep], i)
			inDegree[i]++
		}
	}

	queue := list.New()
	for i := 0; i < n; i++ {
		if inDegree[i] == 0 {
			queue.PushBack(i)
		}
	}

	completionTime := make([]int64, n)

	for queue.Len() > 0 {
		element := queue.Front()
		processIndex := element.Value.(int)
		queue.Remove(element)

		completionTime[processIndex] += executionTime[processIndex]

		for _, neighbor := range graph[processIndex] {
			inDegree[neighbor]--
			if inDegree[neighbor] == 0 {
				completionTime[neighbor] = completionTime[processIndex]
				queue.PushBack(neighbor)
			}
		}
	}

	minTime := int64(0)
	for _, time := range completionTime {
		if time > minTime {
			minTime = time
		}
	}

	fmt.Println(minTime)
}
