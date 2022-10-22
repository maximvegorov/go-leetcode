package keysandrooms

func canVisitAllRooms(rooms [][]int) bool {
	s := []int{0}
	visited := make(map[int]bool)
	for len(s) > 0 {
		r := s[len(s)-1]
		s = s[:len(s)-1]
		if _, ok := visited[r]; ok {
			continue
		}
		visited[r] = true
		for _, k := range rooms[r] {
			if _, ok := visited[k]; ok {
				continue
			}
			s = append(s, k)
		}
	}
	return len(visited) == len(rooms)
}
