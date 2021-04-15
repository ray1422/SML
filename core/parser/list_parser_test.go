package parser

import (
	"testing"

	"github.com/ray1422/SML/core/container"
)

func TestReadLine(t *testing.T) {
	s := `
hi
hello
test
`
	s2 := s
	ans :=
		[]struct {
			a int
			b string
		}{{a: 3, b: "hi"}, {a: 6, b: "hello"}, {a: 5, b: "test"}, {a: 1, b: ""}}

	idx := 0
	total := 0
	for res, i := readLine(&s); i != 0; res, i = readLine(&s) {
		if res != ans[idx].b || i != ans[idx].a {
			t.Error(res, i, "should be", ans[idx].b, ans[idx].a)
		}
		idx++
		total += i
	}
	if idx != len(ans) {
		t.Error("len not correct")
	}
	if total != len(s2) {
		t.Error("str len not correct, expect", len(s2), "output", total)
	}
}

func Test_List(t *testing.T) {
	RegAll()
	content := `
		- list A 1
		- List A 2
		- List B 1
		- List B 2
			- List C 1
			- List C 2
		- List B 3
		- List B 4
1. test 1
2. test 2 **BOLD__ITALIC__**
3. test 3


abc 1. www 2. www
`
	blk := Parse(content)
	container.Dump(blk, 0)
}
