package node_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gellel/binarygo/node"
)

func create(value int) node.Node {
	return node.New(value)
}

func isinstance(property interface{}) bool {
	return reflect.TypeOf(property) == reflect.TypeOf(node.Node{})
}

func TestCreateNode(t *testing.T) {

	fmt.Println("Running test 'CreateNode'.")

	if isinstance(create(5)) == false {
		t.Fatalf("Fatal. Cannot instantiate new Node.")
	}
}

func TestCreateAndAccessValue(t *testing.T) {

	fmt.Println("Running test 'CreateAndAccessValue'.")

	n := create(5)

	if isinstance(n) == false {
		t.Fatalf("Fatal. Cannot instantiate new Node.")
	}
	if reflect.TypeOf(n.Value).Kind() != reflect.Int {
		t.Fatalf("Fatal. Node.Value was assigned a non int type.")
	}
	if n.Value != 5 {
		t.Errorf("Error. Node.Value should be '%d' but is '%d'.", 5, n.Value)
	}
}
