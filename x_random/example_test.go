package x_random_test

import (
	"fmt"
	"github.com/guolei19850528/go_x_pkgs/random"
)

func ExampleRandIntN() {
	n := x_random.RandIntN(12)
	fmt.Println(n)
}

func ExampleRandStrN() {
	str, err := x_random.RandStrN(8, x_random.EN_NORMAL_STR)
	if err == nil {
		fmt.Println(str)
	}
}
