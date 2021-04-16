package parser

import (
	"fmt"
	"testing"
)

func Test_ParseAttr(t *testing.T) {
	s := `{
		a: "yet another string",
		"b":      ww  "  w"w
	}  `

	fmt.Println(len(s))
	a, n := parseAttr(s)
	a.Dump(0)
	fmt.Println(n)

}
