package serializeanddeserializebinarytree

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/maximvegorov/go-leetcode"
)

type TreeNode = leetcode.TreeNode

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return ""
	}
	var buffer strings.Builder
	doSerialize(&buffer, root)
	return buffer.String()
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	if len(data) == 0 {
		return nil
	}
	reader := treeReader{buffer: data}
	return reader.readNode()
}

func doSerialize(buffer *strings.Builder, root *TreeNode) {
	if root == nil {
		return
	}
	buffer.WriteString(strconv.Itoa(root.Val))
	buffer.WriteByte('(')
	if root.Left != nil {
		doSerialize(buffer, root.Left)
	}
	buffer.WriteByte(')')
	buffer.WriteByte('(')
	if root.Right != nil {
		doSerialize(buffer, root.Right)
	}
	buffer.WriteByte(')')
}

type treeReader struct {
	buffer  string
	current int
}

func (r *treeReader) readNode() *TreeNode {
	val, ok := r.readVal()
	if !ok {
		return nil
	}
	r.match('(')
	left := r.readNode()
	r.match(')')
	r.match('(')
	right := r.readNode()
	r.match(')')
	return &TreeNode{Val: val, Left: left, Right: right}
}

func (r *treeReader) match(c byte) {
	if r.buffer[r.current] != c {
		panic(fmt.Errorf("expected %c", c))
	}
	r.current++
}

func (r *treeReader) readVal() (int, bool) {
	start := r.current
	if r.buffer[start] == ')' {
		return 0, false
	}
	if r.buffer[r.current] == '-' {
		r.current++
	}
	for r.current < len(r.buffer) && isDigit(r.buffer[r.current]) {
		r.current++
	}
	res, err := strconv.Atoi(r.buffer[start:r.current])
	if err != nil {
		panic(errors.New("expected val"))
	}
	return res, true
}

func isDigit(c byte) bool {
	return '0' <= c && c <= '9'
}
