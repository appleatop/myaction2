package test2_test

import (
	"fmt"
	"myaction/test2"
	"testing"
	"github.com/bobbae/q"
)

func init() {
	q.O = "stderr"
	q.P = ".*"
}


func TestProgram2(t *testing.T) {
	fmt.Println("testing testprogram 2")
	q.Q("testing testprogram 2")
 	test2.Examplecode()
}
