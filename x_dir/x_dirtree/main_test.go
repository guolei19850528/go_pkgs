package x_dirtree_test

import (
	"fmt"
	"github.com/guolei19850528/go_x_pkgs/x_dir/x_dirtree"
	"testing"
)

func TestTree(t *testing.T) {
	name := "/Users/guolei/Dev/JetBrains/GoLand/oc-projects/oc-go-images/Air/input"
	matchExt := map[string]string{
		".png": ".png",
	}
	dirTree := x_dirtree.Tree(name, matchExt, nil)
	for _, node := range dirTree {
		fmt.Println(node.Children)
	}
}
