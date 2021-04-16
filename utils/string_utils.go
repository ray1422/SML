package utils

var (
	escapeTable = map[string]string{
		"a":  "\a",
		"b":  "\b",
		"f":  "\f",
		"n":  "\n",
		"r":  "\r",
		"t":  "\t",
		"v":  "\v",
		"\\": "\\",
		"'":  "'",
		"\"": "\"",
	}
)

func DeEsc(s string) string {
	if v, ok := escapeTable[s]; ok {
		return v
	}
	return s
}
