package utils

import (
	"fmt"
	"testing"
)

func TestDict(t *testing.T) {
	a := Dict{
		"a": "wwwwwwwwwwww",
		"d": "wwwwwwwwwwww",
		"child": Dict{
			"aa": 1,
			"bb": 1,
		},
	}
	for u := range a.Values() {
		fmt.Println(u)
	}
}
