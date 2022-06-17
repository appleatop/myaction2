package test2_test

import (
	"fmt"
	"myaction2/test2"
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
}

func TestExampleCode(t *testing.T) {
	fmt.Println("testing testprogram 2/ example code")
	q.Q("testing testprogram 2/ example code ")
	output := test2.Examplecode("test pass")
	//output := test2.Examplecode("test error")
	expect := "test2/examplecode: test pass"
	if output != expect {
		t.Errorf("The result (%s) does not match %s", output, expect)
	} else {
		q.Q("run successfully - %s", expect)
	}

}
