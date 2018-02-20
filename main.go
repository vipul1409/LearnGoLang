// Fun with golang. Courtesy rain forest !!
package main

import (
	"fmt"
	"github.com/vipul1409/LearnGoLang/bst"
	"github.com/vipul1409/LearnGoLang/set"
)

func main() {
	demotBst()
	demoSet()
}

func demotBst() {
	root := bst.BuildTree(10, 5)
	bst.PrintInOrder(root)
}

func demoSet() {
	s := set.New()
	s.Add(1, 2, 3, 4, 5)
	fmt.Printf("%v", s)
	s.Remove(3, 5)
	fmt.Printf("%v", s)
}
