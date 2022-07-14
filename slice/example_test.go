package slice_test

import (
	"fmt"
	"github.com/guolei19850528/go_pkgs/slice"
)

func ExampleReverseNoSort() {
	x := []int{1, 4, 7, 6, 3}
	slice.ReverseNoSort(x)
	fmt.Println(x)
	//Output:
	//[3 6 7 4 1]
}
func ExampleReverseNoSortClone() {
	x := []int{1, 4, 7, 6, 3}
	y := slice.ReverseNoSortClone(x)
	fmt.Println("before x ", x)
	fmt.Println("before y ", y)
	y[0] = 9
	fmt.Println("after x ", x)
	fmt.Println("after y ", y)
	//Output:
	//before x [1 4 7 6 3]
	//before y [3 6 7 4 1]
	//after x [1 4 7 6 3]
	//after y [9 6 7 4 1]
}

func ExampleCartesianProduct() {
	var data = [][]int{
		[]int{1, 2, 3},
		[]int{4, 5, 6},
		[]int{7, 8, 9},
	}
	results := slice.CartesianProduct(data...)
	fmt.Println(results)
	//Output:
	//[
	//	[1 4 7],
	//	[1 4 8],
	//	[1 4 9],
	//	[1 5 7],
	//	[1 5 8],
	//	[1 5 9],
	//	[1 6 7],
	//	[1 6 8],
	//	[1 6 9],
	//	[2 4 7],
	//	[2 4 8],
	//	[2 4 9],
	//	[2 5 7],
	//	[2 5 8],
	//	[2 5 9],
	//	[2 6 7],
	//	[2 6 8],
	//	[2 6 9],
	//	[3 4 7],
	//	[3 4 8],
	//	[3 4 9],
	//	[3 5 7],
	//	[3 5 8],
	//	[3 5 9],
	//	[3 6 7],
	//	[3 6 8],
	//	[3 6 9],
	//]
}

func ExampleRand() {
	x := []int{1, 4, 7, 6, 3}
	fmt.Println(slice.Rand(x))
}
