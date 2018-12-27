package node_test

import "testing"
import "reflect"
import "github.com/gellel/binarygo/node"

func CreateNodeTest(t *testing.T) {

	n := node.New(5)

	if reflect.TypeOf(n) != reflect.TypeOf(node.Node{}) {
		t.Fatalf("Cannot create new Node structs.")
	}
}
