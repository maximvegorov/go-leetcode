package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type QNode struct {
	Val         bool
	IsLeaf      bool
	TopLeft     *QNode
	TopRight    *QNode
	BottomLeft  *QNode
	BottomRight *QNode
}

type GNode struct {
	Val       int
	Neighbors []*GNode
}
