// Package node provides the building blocks for creating the binary search tree.
// Nodes are required to distribute N child nodes across the Tree based on the value sum.
// New nodes can be instantiated by using the node.New method.
package node

// Node expresses the children (or leaves) of the binary search tree.
// Each Node contains three properties, Node.left, Node.right and Node.value.
// Node.left and Node.right can be nil during instantiation. Node.value must be assigned.
type Node struct {
	left  *Node
	right *Node
	value int
}

// addLeftNode manages assigning a Left Child Node to the Parent Node.
// Left Child Nodes always contain a value that is smaller than the Parent Node and the Right Child Node.
// Method returns the newly added Left Child Node and the associated Parent Node.
func (parent *Node) addLeftNode(child *Node) (Node, Node) {
	// assumes this has been correctly balanced from calling method.
	parent.left = child
	// @return (Node, Node)
	return *child, *parent
}

// addRightNode manages assigning a Right Child Node to the Parent Node.
// Right Child Nodes always contain a value that is greater than the Parent Node and the Left Child Node.
// Method returns the newly added Right Child Node and the associated Parent Node.
func (parent *Node) addRightNode(child *Node) (Node, Node) {
	// assumes this has been correctly balanced from calling method.
	parent.right = child
	// @return (Node, Node)
	return *child, *parent
}
