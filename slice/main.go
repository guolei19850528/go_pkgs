package slice

import (
	"sort"
)

// RevNoSort
//
//reverse slice or array and no sort
//
// param: x slice or array
func RevNoSort(x any) {
	sort.Slice(x, func(i, j int) bool {
		return i > j
	})
}
