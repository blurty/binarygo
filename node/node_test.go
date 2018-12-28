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
	return reflect.TypeOf(property) == reflect.TypeOf((*node.Node)(nil))
}

func Test(t *testing.T) {

	var n node.Node = create(2)

	if !isinstance(&n) {
		t.Fatalf("Fatal. Cannot create a Parent Node.")
	}
	if !(reflect.TypeOf(n.Value).Kind() == reflect.Int) {
		t.Fatalf("Fatal. Parent Node Value is not numeric.")
	}

	var left, _ node.Node = n.Add(1)

	if !isinstance(&left) {
		t.Fatalf("Fatal. Cannot assign Parent Node a Left Child Node.")
	}

	var right, _ node.Node = n.Add(3)

	if !isinstance(&right) {
		t.Fatalf("Fatal. Cannot assign Parent a Right Child Node.")
	}

	if left.Value > right.Value {
		t.Fatalf("Fatal. Parent Node incorrectly balanced Child Node. Left Child Node is not a lower sum. Left '%d'. Right '%d'.", left.Value, right.Value)
	}

	_, v := n.Visit()

	fmt.Println(v)

	//if v == n.Left
}
