package tools

import (
	"fmt"
	"testing"
)

func TestCompareNumberStrings(t *testing.T) {

	a := "0.00"
	b := "4000"
	value, err := CompareNumberStrings(a, b)
	if err != nil {

		t.Error(err.Error())
	}

	fmt.Println(value)
}
