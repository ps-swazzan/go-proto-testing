package main

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
	//"google.golang.org/protobuf/proto"
	"github.com/stretchr/testify/assert"
	"swazzan/myapp/tutorialpb"
)

func TestProtobufEqual(t *testing.T) {
	var person1 *tutorialpb.Person = &tutorialpb.Person{
		Id:    1,
		Name:  "Saeid",
		Email: "said.el-wazzan@ascent.io",
	}

	var person2 *tutorialpb.Person = &tutorialpb.Person{
		Id:    2,
		Name:  "Saeid",
		Email: "said.el-wazzan@ascent.io",
	}

	t.Logf("Person1: %s", spew.Sdump(person1))

	assert.Equal(t, person1, person2, "The two persons should be equal")

	// The right way:
	// if !proto.Equal(person1, person2) {
	// 	t.Errorf("The two persons are not equal")
	// }
}
