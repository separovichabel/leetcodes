package solution

func AddBinary(a string, b string) string {
	if a < b {
		a, b = b, a
	}

	temp := make([]rune, len(a))

	var carring uint8 = 0
	for aID, bID := len(a)-1, len(b)-1; aID >= 0; {

		if a[aID] == '1' {
			carring++
		}

		if bID >= 0 && b[bID] == '1' {
			carring++
		}

		switch carring {
		case 1:
			temp[aID], carring = '1', 0
		case 2:
			temp[aID], carring = '0', 1
		case 3:
			temp[aID], carring = '1', 1
		}

		aID--
		bID--
	}

	if carring == 1 {
		resposta := make([]rune, len(temp)+1)
		copy(resposta[1:], temp)
		resposta[0] = '1'
		return string(resposta)
	}

	return string(temp)

}
