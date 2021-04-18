package parser

import (
	"testing"

	"github.com/ray1422/SML/core/container"
)

func TestImgParser(t *testing.T) {
	RegAll()
	txt := `!['alt text]']('src://wwwwwwwwwwww' "Title"){float: left}`
	a := ParseNonAsync(txt)
	container.Dump(a, 0)
	txt = `![alt text](src://wwwwwwwwwwww Title **BOLD __ITALIC__ ** wwwww)`
	a = ParseNonAsync(txt)

	container.Dump(a, 0)
}
