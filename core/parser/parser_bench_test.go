package parser

import (
	"testing"
)

func Benchmark_Parser(t *testing.B) {
	RegAll()
	txt := `# Title **BOLD**
	**BOLD \\ \* ESCAPE**
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
	__`
	txt = txt + txt + txt + txt + txt + txt + txt + txt + txt + txt
	for i := 0; i < t.N; i++ {
		Parse(txt)
	}

}
func Benchmark_ParserAsync(t *testing.B) {
	RegAll()
	txt := `# Title **BOLD**
	**BOLD \\ \* ESCAPE**
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
	__`
	txt = txt + txt + txt + txt + txt + txt + txt + txt + txt + txt
	for i := 0; i < t.N; i++ {
		ParseAsync(txt)
	}

}
