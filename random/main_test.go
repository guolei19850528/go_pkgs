package random_test

import (
	"fmt"
	"github.com/guolei19850528/go_pkgs/random"
	"testing"
)

func TestRandIntN(t *testing.T) {
	n := random.RandIntN(12)
	fmt.Println(n)
}

func TestRandStrN(t *testing.T) {
	str, err := random.RandStrN(8, random.EN_NORMAL_STR)
	if err == nil {
		fmt.Println(str)
	}
}
