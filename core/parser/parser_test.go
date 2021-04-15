package parser

import (
	"testing"

	"github.com/ray1422/SML/core/container"
)

func Test_TMP(t *testing.T) {
	RegAll()
	a := Parse(`
- list
`)
	container.Dump(a, 0)
}

func Test_Parser(t *testing.T) {
	RegAll()
	txt := `# Title **BOLD**
**BOLD \\ \* ESCAPE**

- test
- test2
	- test3


\\n
~~__ITALIC**BOLD**__ LOL~~
# Title
` + "```cpp" + `
#include <bits/stdc++.h>
using namespace std;
int main(void) {
	cout << "Hello World!";
}
` + "```" + `
__LOL__
__
` + "`int main(){}`{:.cpp}\n"

	a := ParseNonAsync(txt)
	container.Dump(a, 0)
}
