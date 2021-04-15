package utils

func StrNCmp(a string, b string, n int) bool {
	if len(a) < n || len(b) < n {
		return false
	}
	if a[:n] == b[:n] {
		return true
	}
	return false
}
