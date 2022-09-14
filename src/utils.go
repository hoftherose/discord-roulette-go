package util

func JoinStrings(sep string, str ...string) string {
	if len(str) == 0 {
		return ""
	}
	resp := str[0]
	for _, s := range str[1:] {
		resp += sep + s
	}
	return resp
}
