package main

import (
	"fmt"
	"myaction/test1"
	"myaction/test2"
	"myaction/test3" 
	"myaction/test4/a"
	t3 "myaction/test5/a/b/c/test3"
	"github.com/bobbae/q"
)

func init() {
	q.O = "stderr"
	q.P = ".*"
}

func main() {
	q.Q("TESTTESTTESTTEST TESTGO3")
	fmt.Println("mainprogram")
	test1.Tmain1()
	test2.Examplecode()
	test3.Tmain3()
	t3.Tmain3Special()
	a.Tmain4()
	q.Q("test")
}
