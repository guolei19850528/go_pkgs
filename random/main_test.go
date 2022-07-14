package random_test

import (
	"fmt"
	"github.com/guolei19850528/go_pkgs/random"
	"testing"
)

func TestRandStr(t *testing.T) {
	str, err := random.RandStr(8, random.EN_NORMAL_STR)
	if err == nil {
		fmt.Println(str)
	}
}
