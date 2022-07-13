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
// param: args slice
// return:
func CartesianProduct(args ...[]any) []any {
	results := make([]any, 0, 0)
	firstArgs, otherArgs := args[0], args[1:]
	for _, firstArg := range firstArgs {
		groupBase := append(make([]any, 0, 0), firstArg)
		groupData := make([]any, 0, 0)
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
func cartesianProductHandler(groupBase *[]any, groupData *[]any, args ...[]any) {
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
