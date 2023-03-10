package solution

func AddBinary(a string, b string) string {
	if len(a) < len(b) {
		a, b = b, a
	}

	temp := make([]byte, len(a))

	var carring uint8 = 0
	for aID, bID := len(a)-1, len(b)-1; aID >= 0; {

		if a[aID] == '1' {
			carring++
		}

		if bID >= 0 && b[bID] == '1' {
			carring++
		}

		switch carring {
		case 0:
			temp[aID], carring = '0', 0
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
		resposta := make([]byte, len(temp)+1)
		copy(resposta[1:], temp)
		resposta[0] = '1'
		return string(resposta)
	}

	return string(temp)

}

func AddBinary2(a string, b string) string {
	if len(a) < len(b) {
		a, b = b, a
	}

	temp := ""

	var carring uint8 = 0

	for aID, bID := len(a)-1, len(b)-1; aID >= 0; {

		if a[aID] == '1' {
			carring++
		}

		if bID >= 0 && b[bID] == '1' {
			carring++
		}

		switch carring {
		case 0:
			temp, carring = "0"+temp, 0
		case 1:
			temp, carring = "1"+temp, 0
		case 2:
			temp, carring = "0"+temp, 1
		case 3:
			temp, carring = "1"+temp, 1
		}

		aID--
		bID--
	}

	if carring == 1 {
		return string("1" + temp)
	}

	return string(temp)

}
