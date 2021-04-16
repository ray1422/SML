package utils

func StrIdxSafe(s string, i int) bool {
	return len(s) > i
}

func QuoteRM(s string) string {
	if len(s) < 2 { // "" or "?"
		return s
	}
	if (s[0] == '"' && s[len(s)-1] == '"') || (s[0] == '\'' && s[len(s)-1] == '\'') {
		s = s[1 : len(s)-1]
	}
	return s
}
