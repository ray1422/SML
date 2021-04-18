package parser

import (
	"testing"
)

const (
	txt4test = `# Title **BOLD**
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
![image](src:/// alt){wwwwww: test}
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
)

func BenchmarkParser(t *testing.B) {
	RegAll()
	txt := txt4test
	txt = txt + txt + txt + txt + txt + txt + txt + txt + txt + txt
	for i := 0; i < t.N; i++ {
		Parse(txt)
	}
}
func BenchmarkParserAsync(t *testing.B) {
	RegAll()
	txt := txt4test
	txt = txt + txt + txt + txt + txt + txt + txt + txt + txt + txt
	for i := 0; i < t.N; i++ {
		ParseAsync(txt)
	}
}
