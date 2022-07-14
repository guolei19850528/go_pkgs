package random

import (
	"fmt"
	"math"
	"math/rand"
	"reflect"
	"time"
)

const (
	EN_LOWER_STR  = "abcdefghijklmnopqrstuvwxyz"
	EN_UPPER_STR  = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	EN_SYMBOL_STR = "`~!@#$%^&*()-=_+[]{}|"
	NUMBER_STR    = "0123456789"
	EN_NORMAL_STR = EN_LOWER_STR + EN_UPPER_STR + NUMBER_STR
)

// RandStrN
//
// param: n
// param: chars
// return:
func RandStrN(n int, chars string) (string, error) {
	if n < 0 {
		n = int(math.Abs(float64(n)))
	}
	if reflect.TypeOf(chars).Kind() != reflect.String || len(chars) < 0 {
		return "", fmt.Errorf("%s is not string or empty", chars)
	}

	bytes := make([]byte, 0, 0)
	for i := 0; i < n; i++ {
		bytes = append(bytes, chars[RandIntN(len(chars))])
	}
	return string(bytes), nil
}

// RandIntN
//
// param: n
// return:
func RandIntN(n int) int {
	rand.Seed(time.Now().UnixMicro())
	return rand.Intn(n)
}
