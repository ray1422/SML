package parser

import (
	"fmt"
	"testing"
)

func Benchmark_Parser(t *testing.B) {
	RegAll()
	txt := `# Title **BOLD**
	**BOLD \\ \* ESCAPE**
	- list
		- list
		- list
		- list
			1. wwww
			2. wwwww
			4. wwwww
	- list
1. wwww
2. eeeee
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
	fmt.Printf("len: %d\n", len(txt))
	for i := 0; i < t.N; i++ {
		ParseAsync(txt)
	}

}
