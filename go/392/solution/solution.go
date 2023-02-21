package solution

func IsSubsequence(s string, t string) bool {
	sIndex, tIndex := 0, 0

	for ; tIndex < len(t); tIndex++ {
		if sIndex+1 > len(s) {
			return true
		}
		if s[sIndex] == t[tIndex] {
			sIndex++
		}
	}

	if sIndex+1 > len(s) {
		return true
	}

	return false

}
