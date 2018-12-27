// Package node provides the building blocks for creating the binary search tree.
// Nodes are required to distribute N child nodes across the Tree based on the value sum.
// New nodes can be instantiated by using the node.New method.
package node

// Node expresses the children (or leaves) of the binary search tree.
// Each Node contains three properties, Node.left, Node.right and Node.value.
// Node.left and Node.right can be nil during instantiation. Node.value must be assigned.
type Node struct {
	left  *Node // Lower value Node.
	right *Node // Greater value Node.
	value int   // Sum of current Node.
}

// New instantiates a new Node struct.
// Constructor expects the value to be an integer.
// TODO: Allow overloading and search from this data type?
func New(value int) Node {
	return Node{value: value}
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

// determineSide confirms what position the provided Child Node is to be assigned.
// Operation works on checking whether the value of the Child Node is greater than the Parent Node.
// Should the value be smaller, a left assignment operation is performed. If greater, the right assignment occurs.
// Method returns the side the Child Node was assigned to.
func (parent *Node) determineSide(child *Node) (Node, Node) {
	if child.value < parent.value {
		if parent.left == nil {
			return parent.addLeftNode(child)
		}
		return parent.left.determineSide(child)
	} else if child.value > parent.value {
		if parent.right == nil {
			return parent.right.addRightNode(child)
		}
		return parent.right.determineSide(child)
	}
	return *child, *parent
}
