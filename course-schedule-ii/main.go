package main

func findOrder(numCourses int, prerequisites [][]int) []int {
	dependencies := make(map[int][]int)
	invDependencies := make(map[int][]int)
	for _, p := range prerequisites {
		dependencies[p[0]] = append(dependencies[p[0]], p[1])
		invDependencies[p[1]] = append(invDependencies[p[1]], p[0])
	}
	ends := []int{}
	for c := 0; c < numCourses; c++ {
		if len(invDependencies[c]) == 0 {
			ends = append(ends, c)
		}
	}
	if len(ends) == 0 {
		return []int{}
	}
	weights := make([]int, numCourses)
	weight := 1
	for _, end := range ends {
		q := []int{end}
		inProgress := make(map[int]struct{})
		for len(q) > 0 {
			n := q[len(q)-1]
			if weights[n] > 0 {
				q = q[:len(q)-1]
				continue
			}
			inProgress[n] = struct{}{}
			ready := true
			for _, c := range dependencies[n] {
				if weights[c] == 0 {
					if _, found := inProgress[c]; found {
						return []int{}
					}
					if ready {
						ready = false
					}
					q = append(q, c)
				}
			}
			if ready && weights[n] == 0 {
				weights[n] = weight
				weight++
				q = q[:len(q)-1]
				delete(inProgress, n)
			}
		}
	}
	schedule := make([]int, numCourses)
	for c := 0; c < numCourses; c++ {
		if weights[c] == 0 {
			return []int{}
		}
		schedule[weights[c]-1] = c
	}
	return schedule
}
