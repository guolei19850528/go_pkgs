package random_test

import (
	"fmt"
	"github.com/guolei19850528/go_pkgs/random"
)

func ExampleRandIntN() {
	n := random.RandIntN(12)
	fmt.Println(n)
}

func ExampleRandStrN() {
	str, err := random.RandStrN(8, random.EN_NORMAL_STR)
	if err == nil {
		fmt.Println(str)
	}
}
