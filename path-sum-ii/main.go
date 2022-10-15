package path_sum_ii

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, targetSum int) [][]int {
	if root == nil {
		return [][]int{}
	}
	finder := pathFinder{remain: targetSum}
	finder.walk(root)
	return finder.paths
}

type pathFinder struct {
	remain int
	path   []int
	paths  [][]int
}

func (f *pathFinder) walk(root *TreeNode) {
	if f.isLeaf(root) {
		if f.remain == root.Val {
			path := make([]int, len(f.path)+1)
			copy(path, f.path)
			path[len(f.path)] = root.Val
			f.paths = append(f.paths, path)
		}
		return
	}
	if root.Left != nil {
		f.pushNode(root)
		f.walk(root.Left)
		f.popNode(root)
	}
	if root.Right != nil {
		f.pushNode(root)
		f.walk(root.Right)
		f.popNode(root)
	}
}

func (f *pathFinder) pushNode(node *TreeNode) {
	f.path = append(f.path, node.Val)
	f.remain -= node.Val
}

func (f *pathFinder) popNode(node *TreeNode) {
	f.remain += node.Val
	f.path = f.path[:len(f.path)-1]
}
func (f *pathFinder) isLeaf(node *TreeNode) bool {
	return node.Left == nil && node.Right == nil
}
