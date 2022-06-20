package test3_test

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

func TestProgram3(t *testing.T) {
	fmt.Println("testing testprogram 3")
	q.Q("testing testprogram 3")
}

func TestRandom1(t *testing.T) {
	fmt.Println("testing testprogram 3/ random1")
	q.Q("testing testprogram 3/ random1")
	//output := test2.Examplecode("test pass")
	//output := test2.Examplecode("test error")
	output := test2.Examplecode("test again for program3")
	expect := "test2/examplecode: test again for program3"
	if output != expect {
		t.Errorf("The result (%s) does not match %s", output, expect)
	} else {
		q.Q("run successfully: ", expect)

	}

}

func TestRandom2(t *testing.T) {
	fmt.Println("TestRandom2")
}

func TestMain3Tmp(t *testing.T) {
	fmt.Println("TestMain3Tmp")
}
