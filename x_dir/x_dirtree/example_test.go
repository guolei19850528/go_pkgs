package x_dirtree_test

import (
	"fmt"
	"github.com/guolei19850528/go_x_pkgs/x_dir/x_dirtree"
	"sort"
	"strconv"
)

func ExampleTree() {
	name := "/Users/yourname/images"
	matchExt := map[string]string{
		".png": ".png",
	}
	sortFunc := func(dirTreeNodes []*x_dirtree.TreeNode) {
		sort.Slice(dirTreeNodes, func(i, j int) bool {
			sortKey1, err1 := strconv.ParseInt(dirTreeNodes[i].BaseName, 10, 32)
			sortKey2, err2 := strconv.ParseInt(dirTreeNodes[j].BaseName, 10, 32)
			if err1 == nil && err2 == nil {
				return sortKey1 < sortKey2
			}
			return dirTreeNodes[i].BaseName < dirTreeNodes[j].BaseName
		})
	}

	dirTree := x_dirtree.Tree(name, matchExt, sortFunc)
	fmt.Println(dirTree)
}
