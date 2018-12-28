// Package binarygo creates a binary search tree for Gopher.
// Binarygo can create a search tree using pointers or values.
// Use Tree.Pointer for pointers. Use Tree.Value for values.
package binarygo

import "fmt"
import "github.com/gellel/binarygo/node"

func main() {
	fmt.Println("binarygo is running.")

	n := node.New(10)

	fmt.Println(n)
}
