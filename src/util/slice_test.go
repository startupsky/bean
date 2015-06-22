package util

import (
	"fmt"
	"testing"
)

func Test_remove(t *testing.T) {
	a := BeanSlice{1, 2, 3, 4, 5, 6}
	fmt.Println(a.Remove(3))
}
