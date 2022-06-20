package a_test

import (
	"fmt"
	"myaction2/test4/a"
	"testing"

	"github.com/bobbae/q"
)

func init() {
	q.O = "stderr"
	q.P = ".*"
}

func TestProgram4(t *testing.T) {
	fmt.Println("testing testprogram 4")
	q.Q("testing testprogram 4")
}

func TestMain4Tmp(t *testing.T) {
	fmt.Println("TestMain4Tmp")
	a.Tmain4()
	fmt.Println("Should see the testmain4 line")
}
