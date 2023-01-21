package verifypreorderserializationofabinarytree

import "strings"

func isValidSerialization(preorder string) bool {
	if len(preorder) == 0 || preorder == "#" {
		return true
	}
	nodes := strings.Split(preorder, ",")
	cnt, valid := isValid(nodes)
	if !valid {
		return false
	}
	return len(nodes) == cnt
}

func isValid(nodes []string) (int, bool) {
	if nodes[0] == "#" || len(nodes) < 3 {
		return 0, false
	}
	cnt := 1
	for i := 1; i <= 2; i++ {
		if nodes[cnt] == "#" {
			cnt += 1
		} else {
			childCnt, valid := isValid(nodes[cnt:])
			if !valid {
				return 0, false
			}
			cnt += childCnt
		}
	}
	return cnt, true
}
