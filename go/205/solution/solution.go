package solution

func IsIsomorphic(s string, t string) bool {
	sMatches := map[byte]byte{}
	tMatches := map[byte]byte{}

	for i := 0; i < len(s); i++ {
		sMatch, sExists := sMatches[s[i]]
		tMatch, tExists := tMatches[t[i]]
		if sExists && tExists && (sMatch == t[i] && tMatch == s[i]) {
			continue
		}

		if !sExists && !tExists {
			sMatches[s[i]] = t[i]
			tMatches[t[i]] = s[i]
			continue
		}

		return false
	}
	return true
}
