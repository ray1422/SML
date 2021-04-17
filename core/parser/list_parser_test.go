package parser

import (
	"testing"

	"github.com/ray1422/SML/core/container"
)

func TestList(t *testing.T) {
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
1. 	test 1
2. 	test 2 **BOLD__ITALIC__**
30.	test 3


abc 


1.			www
2.       www
`
	blk := Parse(content)
	container.Dump(blk, 0)
}

func TestOList(t *testing.T) {
	RegAll()
	content := `    1. testA
      2. testB
     3. testC
wwwwwwwwww
		1. test
		2. test
1. A
2. B
	1. BA
	2. BB
	3. BC
		1. BCA
		2. BCA
		3. WWW
3. WWW
`
	blk := Parse(content)
	container.Dump(blk, 0)
}
