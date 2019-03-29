package tree

import "fmt"

type TreeNode struct {
	Data interface{}
	Left *TreeNode
	Right *TreeNode
}

type BinTree interface {
	InorderTraversal()
	PreorderTraversal()
	PostorderTraversal()
	LevelorderTraversal()
	GetHeight() int
}

func NewBinTree() *TreeNode {
	return &TreeNode{Data:nil,Left:nil,Right:nil}
}

func (tree *TreeNode) InorderTraversal() {
	if tree != nil {
		tree.Left.InorderTraversal()
		fmt.Println(tree.Data)
		tree.Right.InorderTraversal()
	}
}

func (tree *TreeNode) PreorderTraversal() {
	if tree != nil {
		fmt.Println(tree.Data)
		tree.Left.PreorderTraversal()
		tree.Right.PreorderTraversal()
	}
}

func (tree *TreeNode) PostorderTraversal() {
	if tree != nil {
		tree.Left.PostorderTraversal()
		tree.Right.PostorderTraversal()
		fmt.Println(tree.Data)
	}
}

func (tree *TreeNode) LevelorderTraversal() {
	q := make([]*TreeNode,0)
	q = append(q,tree)
	for len(q)!=0 {
		element := q[0]
		q = q[1:]
		fmt.Println(element.Data)
		if element.Left != nil {
			q = append(q,element.Left)
		}
		if element.Right != nil {
			q = append(q,element.Right)
		}
	}
}

func (tree *TreeNode) GetHeight() (res int) {
	var HL,HR int

	if tree != nil {
		HL = tree.Left.GetHeight()
		HR = tree.Right.GetHeight()
		res =  HR + 1
		if HL > HR {
			res = HL + 1
		}
		return 
	}
	return 0
}