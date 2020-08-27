package main

import (
	"fmt"
	"learngo/lang/tree"
)

//遍历二叉树
type myTreeNode struct {
	node *tree.Node
}

func (myNode *myTreeNode) postOrder() {
	if myNode.node == nil || myNode == nil {
		return
	}
	left := myTreeNode{myNode.node.Left}
	right := myTreeNode{myNode.node.Right}
	left.postOrder()
	right.postOrder()
	myNode.node.Print()
}

func main() {
	var root tree.Node
	root = tree.Node{Value: 3}
	root.Left = &tree.Node{}
	root.Right = &tree.Node{5, nil, nil}
	root.Right.Left = new(tree.Node)
	root.Left.Right = tree.CreateNode(2)
	root.Right.Left = tree.CreateNode(4)
	fmt.Print("In-order:")
	myRoot := myTreeNode{&root}
	myRoot.postOrder()
	fmt.Println()

}
