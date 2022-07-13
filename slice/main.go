package slice

import (
	"sort"
)

// ReverseNoSort
//
//reverse slice or array and no sort
//
// param: x slice or array
func ReverseNoSort(x any) {
	sort.Slice(x, func(i, j int) bool {
		return i > j
	})
}

// CartesianProduct cartesian permutation and combination
//
// param: args slice or array
// return:
func CartesianProduct[T any](args ...[]T) [][]T {
	results := make([][]T, 0, 0)
	firstArgs, otherArgs := args[0], args[1:]
	for _, firstArg := range firstArgs {
		groupBase := append(make([]T, 0, 0), firstArg)
		groupData := make([][]T, 0, 0)
		cartesianProductHandler(&groupBase, &groupData, otherArgs...)
		results = append(results, groupData...)
	}
	return results
}

// cartesianProductHandler cartesian permutation and combination handler
//
// param: groupRow
// param: groupData
// param: args
func cartesianProductHandler[T any](groupBase *[]T, groupData *[][]T, args ...[]T) {
	firstArgs, otherArgs := args[0], args[1:]
	for _, firstArg := range firstArgs {
		if len(args) > 1 {
			groupBaseTmp := append(*groupBase, firstArg)
			cartesianProductHandler(&groupBaseTmp, groupData, otherArgs...)
		} else {
			*groupData = append(*groupData, append(*groupBase, firstArg))
		}
	}
}
