package tree

import "fmt"

type Node struct {
	Value       int
	Left, Right *Node
}

func (node Node) Print() {
	fmt.Print(node.Value, "")
}

//这些其实就是相当于对象的方法 接收者就是对象本身 在go里叫结构体/其他类型
func (node *Node) setValue(v int) {
	if node == nil {
		fmt.Println("setting value to nil" + "node .Ignored")
		return
	}
	node.Value = v
}

//设置子节点 对比上面的方法
func CreateNode(value int) *Node {
	return &Node{Value: value}
}
