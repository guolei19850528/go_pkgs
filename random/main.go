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

// RandStr
//
// param: length
// param: chars
// return:
func RandStr(length int, chars string) (string, error) {
	if length < 0 {
		length = int(math.Abs(float64(length)))
	}
	if reflect.TypeOf(chars).Kind() != reflect.String || len(chars) < 0 {
		return "", fmt.Errorf("%s is not string or empty", chars)
	}
	rand.Seed(time.Now().UnixMicro())
	bytes := make([]byte, 0, 0)
	for i := 0; i < length; i++ {
		bytes = append(bytes, chars[rand.Intn(len(chars))])
	}
	return string(bytes), nil
}
