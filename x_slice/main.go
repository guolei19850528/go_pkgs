package x_slice

import (
	"github.com/guolei19850528/go_x_pkgs/x_random"
	"sort"
)

// ReverseNoSort
//
//reverse slice or array and no sort
//
// param: x slice or array
func ReverseNoSort[T any](x []T) {
	sort.Slice(x, func(i, j int) bool {
		return i > j
	})
}

// ReverseNoSortClone
//
//reverse slice or array and no sort and clone
//
// param: x slice or array
// return:x clone
func ReverseNoSortClone[T any](x []T) []T {
	cloneX := Clone(x)
	ReverseNoSort(cloneX)
	return cloneX
}

//Clone
//
// param: x slice or array
// return:x clone
func Clone[T any](x []T) []T {
	return append(make([]T, 0, 0), x...)
}

// CartesianProduct cartesian permutation and combination
//
// param: args slice or array
// return:
func CartesianProduct[T any](args ...[]T) [][]T {
	results := make([][]T, 0, 0)
	firstArgs, otherArgs := args[0], args[1:]
	for _, firstArg := range firstArgs {
		base := append(make([]T, 0, 0), firstArg)
		data := make([][]T, 0, 0)
		cartesianProductHandler(&base, &data, otherArgs...)
		results = append(results, data...)
	}
	return results
}

// cartesianProductHandler cartesian permutation and combination handler
//
// param: base
// param: data
// param: args
func cartesianProductHandler[T any](base *[]T, data *[][]T, args ...[]T) {
	firstArgs, otherArgs := args[0], args[1:]
	for _, firstArg := range firstArgs {
		if len(args) > 1 {
			baseTmp := append(*base, firstArg)
			cartesianProductHandler(&baseTmp, data, otherArgs...)
		} else {
			*data = append(*data, append(*base, firstArg))
		}
	}
}

// Rand
//
// param: x
// return:
func Rand[T any](x []T) T {
	return x[x_random.RandIntN(len(x))]
}
