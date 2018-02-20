// Fun with golang. Courtesy rain forest !!
package main

import (
	"github.com/vipul1409/LearnGoLang/bst"
)

func main() {
	root := bst.BuildTree(10, 5)
	bst.PrintInOrder(root)
}
