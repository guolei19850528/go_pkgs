package random_test

import (
	"fmt"
	"github.com/guolei19850528/go_pkgs/random"
)

func ExampleRandStr() {
	str, err := random.RandStr(32, random.EN_NORMAL_STR)
	if err == nil {
		fmt.Println(str)
	}
	//Output:
	//n8XoVL2LIxUfUBBMitAEbWbf0gpey9v0
}
