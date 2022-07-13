package slice_test

import (
	"fmt"
	"github.com/guolei19850528/go_pkgs/slice"
)

func ExampleRevNoSort() {
	x := []int{1, 4, 7, 6, 3}
	slice.RevNoSort(x)
	fmt.Println(x)
	//Output:
	//[3 6 7 4 1]
}
