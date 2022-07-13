package slice_test

import (
	"fmt"
	"github.com/guolei19850528/go_pkgs/slice"
	"reflect"
	"testing"
)

func TestRevNoSort(t *testing.T) {
	x := []int{1, 4, 7, 6, 3}
	slice.RevNoSort(x)
	fmt.Println(x)
}

func TestRevNoSortClone(t *testing.T) {
	x := []int{1, 4, 7, 6, 3}

	typeOfX := reflect.TypeOf(x)
	fmt.Println(typeOfX.Name())
	fmt.Println(typeOfX.Kind())
	fmt.Println(typeOfX.Field(0))

}
