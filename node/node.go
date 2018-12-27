// Package node provides the building blocks for creating the binary search tree.
// Nodes are required to distribute N child nodes across the Tree based on the value sum.
// New nodes can be instantiated by using the node.New method.
package node

import (
	"fmt"
	"reflect"
)

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

// Search attempts to find a set value in the connected Nodes.
func (parent *Node) Search(value int) (Node, bool) {
	if parent.value == value {
		return *parent, true
	}
	if parent.hasLeftNode() || parent.hasRightNode() {
		return *parent, false
	}
	if value < parent.value {
		return parent.left.Search(value)
	}
	if value > parent.value {
		return parent.right.Search(value)
	}
	return *parent, false
}

// Visit will automatically walk through the Child Nodes of the accessed Parent Node.
func (parent *Node) Visit() Node {

	fmt.Println(parent.value)

	if parent.hasLeftNode() {
		return parent.left.Visit()
	} else if parent.hasRightNode() {
		return parent.right.Visit()
	}
	return *parent
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
		if parent.hasLeftNode() {
			return parent.addLeftNode(child)
		}
		return parent.left.determineSide(child)
	} else if child.value > parent.value {
		if parent.hasRightNode() {
			return parent.right.addRightNode(child)
		}
		return parent.right.determineSide(child)
	}
	return *child, *parent
}

// isNode checks whether the provided property is a Node.
func (parent *Node) isNode(property interface{}) bool {
	return reflect.TypeOf(property) == reflect.TypeOf(Node{})
}

// hasLeftSide tests whether the Parent Node has a Node assigned to its left side.
func (parent *Node) hasLeftNode() bool {
	return parent.isNode(parent.left)
}

// hasRightSide tests whether the Parent Node has a Node assigned to its right side.
func (parent *Node) hasRightNode() bool {
	return parent.isNode(parent.right)
}
