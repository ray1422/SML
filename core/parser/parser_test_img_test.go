package parser

import (
	"fmt"
	"testing"

	"github.com/ray1422/SML/core/container"
)

func Test_ImgParser(t *testing.T) {
	RegAll()
	txt := `!['alt text]']('src://wwwwwwwwwwww' "Title"){float: left}`
	a := ParseNonAsync(txt)
	container.Dump(a, 0)
	txt = `![alt text](src://wwwwwwwwwwww Title **BOLD __ITALIC__ ** wwwww)`
	a = ParseNonAsync(txt)
	fmt.Println(a.Children()[0].(*container.ImageBlock))
	container.Dump(a, 0)
}
