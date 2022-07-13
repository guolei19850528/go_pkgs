package convert

import (
	"fmt"
)

// AnyToString any to string
//
// param: x
// return:
func AnyToString(x any) string {
	return fmt.Sprintf("%v", x)
}
