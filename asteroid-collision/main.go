package asteroid_collision

func asteroidCollision(asteroids []int) []int {
	s := make([]int, 0, len(asteroids))
	for _, size := range asteroids {
		if size > 0 || len(s) == 0 || s[len(s)-1] < 0 {
			s = append(s, size)
			continue
		}
		absSize := -size
		for len(s) > 0 && s[len(s)-1] > 0 && s[len(s)-1] < absSize {
			s = s[:len(s)-1]
		}
		if len(s) > 0 {
			top := s[len(s)-1]
			if top == absSize {
				s = s[:len(s)-1]
			} else if top < 0 {
				s = append(s, size)
			}
		} else {
			s = append(s, size)
		}
	}
	return s
}
