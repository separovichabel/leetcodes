package solution

import "strings"

func LongestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	prefix := strs[0]

	for i := 1; i < len(strs); i++ {

		if prefix == "" {
			return ""
		}

		prefix = commonPrefix(strs[i], prefix)
	}

	return prefix
}

func commonPrefix(str, prefix string) string {
	for prefix != "" {
		if strings.HasPrefix(str, prefix) {
			return prefix
		}

		prefix = prefix[:len(prefix)-1]
	}

	return prefix
}
