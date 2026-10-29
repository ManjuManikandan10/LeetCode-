
func basicCalculatorIV(expression string, evalvars []string, evalints []int) []string {
	mapE := make(map[string]string)
	for i := range evalvars {
		mapE[evalvars[i]] = fmt.Sprintf("%d", evalints[i])
	}
	sort.Slice(evalvars, func(i, j int) bool {
		return len(evalvars[i]) > len(evalvars[j])
	})
	expression = strings.ReplaceAll(expression, " ", "")
	mp, val := basicCalculator(expression, mapE)
	rs := []string{}
	for k, v := range mp {
		if v != 0 {
			rs = append(rs, k)
		}
	}
	sort.Slice(rs, func(i, j int) bool {
		ar1 := strings.Split(rs[i], "*")
		ar2 := strings.Split(rs[j], "*")
		if len(ar1) > len(ar2) {
			return true
		}
		if len(ar1) < len(ar2) {
			return false
		}
		for a := range ar1 {
			if ar1[a] < ar2[a] {
				return true
			}
			if ar1[a] > ar2[a] {
				return false
			}
		}
		return true
	})
	for i := range rs {
		rs[i] = fmt.Sprintf("%d*%s", mp[rs[i]], rs[i])
	}
	if val != 0 {
		rs = append(rs, fmt.Sprintf("%d", val))
	}
	return rs

}
func basicCalculator(expression string, mapE map[string]string) (map[string]int, int) {
	i1 := strings.Index(expression, "(")
	if i1 == -1 {
		return getValue(expression, mapE)
	}
	mapVal := make(map[string]int)
	var valOnly int
	arrVal := []string{}
	arrCompute := []byte{}
	for i1 != -1 {
		if i1 > 0 {
			if expression[i1-1] == '*' {
				for j := i1 - 1; j > 0; j-- {
					if expression[j] == '+' || expression[j] == '-' {
						arrVal = append(arrVal, expression[:j])
						arrCompute = append(arrCompute, expression[j])
						expression = expression[j+1:]
						i1 = strings.Index(expression, "(")
						break
					}
				}
				arrVal = append(arrVal, expression[:i1-1])
				arrCompute = append(arrCompute, expression[i1-1])
			} else {
				arrVal = append(arrVal, expression[:i1-1])
				arrCompute = append(arrCompute, expression[i1-1])
			}
			expression = expression[i1:]
		}
		open := 0
		for i := 0; i < len(expression); i++ {
			if expression[i] == '(' {
				open++
			}
			if expression[i] == ')' {
				open--
				if open == 0 {
					arrVal = append(arrVal, expression[1:i])
					if i != len(expression)-1 {
						if expression[i+1] == '+' {
							arrCompute = append(arrCompute, expression[i+1])
							expression = expression[i+2:]
						} else if expression[i+1] == '-' {
							if expression[i+2] == '(' {
								arrCompute = append(arrCompute, expression[i+1])
								expression = expression[i+2:]
							} else {
								arrCompute = append(arrCompute, '+')
								expression = expression[i+1:]
							}

						} else {
							if expression[i+2] == '(' {
								arrCompute = append(arrCompute, expression[i+1])
								expression = expression[i+2:]
							} else {
								arrCompute = append(arrCompute, expression[i+1])
								add := false
								for j := i + 2; j < len(expression); j++ {
									if expression[j] == '(' {
										break
									}
									if expression[j] == '+' {
										arrVal = append(arrVal, expression[i+2:j])
										arrCompute = append(arrCompute, expression[j])
										expression = expression[j+1:]
										add = true
										break
									}
									if expression[j] == '-' {
										if expression[j+1] == '(' {
											arrVal = append(arrVal, expression[i+2:j])
											arrCompute = append(arrCompute, expression[j])
											expression = expression[j+1:]
											add = true
										} else {
											arrVal = append(arrVal, expression[i+2:j])
											arrCompute = append(arrCompute, '+')
											expression = expression[j:]
											add = true
										}
										break
									}
								}
								if !add {
									expression = expression[i+2:]
								}
							}
						}

						i1 = strings.Index(expression, "(")
					} else {
						expression = ""
						i1 = -1
					}
					break
				}
			}
		}
	}

	if len(expression) > 0 {
		arrVal = append(arrVal, expression)
	}

	val := 1
	mapVal0, valOnly0 := basicCalculator(arrVal[0], mapE)
	arrVal = arrVal[1:]
	for len(arrCompute) > 0 {
		switch arrCompute[0] {
		case '+', '-':
			mapVal2, valOnly2 := basicCalculator(arrVal[0], mapE)
			if val == 1 {
				mapVal, valOnly = plus(mapVal, mapVal0, valOnly, valOnly0)
				mapVal0, valOnly0 = mapVal2, valOnly2
			} else {
				mapVal, valOnly = sub(mapVal, mapVal0, valOnly, valOnly0)
				mapVal0, valOnly0 = mapVal2, valOnly2
			}
			val = 1
			if arrCompute[0] == '-' {
				val = -1
			}
		default:
			mapVal2, valOnly2 := basicCalculator(arrVal[0], mapE)
			mapVal0, valOnly0 = mul(mapVal0, mapVal2, valOnly0, valOnly2)
		}
		arrVal = arrVal[1:]
		arrCompute = arrCompute[1:]
	}
	if val == 1 {
		return plus(mapVal, mapVal0, valOnly, valOnly0)
	} else {
		return sub(mapVal, mapVal0, valOnly, valOnly0)
	}
}
func plus(mapVal1, mapVal2 map[string]int, valOnly1, valOnly2 int) (map[string]int, int) {
	for k, v := range mapVal2 {
		mapVal1[k] += v
		if mapVal1[k] == 0 {
			delete(mapVal1, k)
		}
	}
	return mapVal1, valOnly1 + valOnly2
}
func sub(mapVal1, mapVal2 map[string]int, valOnly1, valOnly2 int) (map[string]int, int) {
	for k, v := range mapVal2 {
		mapVal1[k] -= v
		if mapVal1[k] == 0 {
			delete(mapVal1, k)
		}
	}
	return mapVal1, valOnly1 - valOnly2
}
func mul(mapVal1, mapVal2 map[string]int, valOnly1, valOnly2 int) (map[string]int, int) {
	mapRs := make(map[string]int)
	if valOnly1 != 0 {
		for k, v := range mapVal2 {
			mapRs[k] += valOnly1 * v
			if mapRs[k] == 0 {
				delete(mapRs, k)
			}
		}
	}
	if valOnly2 != 0 {
		for k, v := range mapVal1 {
			mapRs[k] += valOnly2 * v
			if mapRs[k] == 0 {
				delete(mapRs, k)
			}
		}
	}
	for k, v := range mapVal1 {
		ar1 := strings.Split(k, "*")
		for k2, v2 := range mapVal2 {
			ar2 := strings.Split(k2, "*")
			ar2 = append(ar2, ar1...)
			v2 *= v
			mapRs[getVal(ar2)] += v2
		}
	}
	return mapRs, valOnly1 * valOnly2
}
func getValue(expression string, mapE map[string]string) (map[string]int, int) {
	mapVal := make(map[string]int)
	var valOnly int64
	arrVal := []string{}
	arrCompute := []byte{}
	i0 := 0
	for i := range expression {
		switch expression[i] {
		case '+':
			arrVal = append(arrVal, expression[i0:i])
			arrCompute = append(arrCompute, '+')
			i0 = i + 1
		case '-':
			if i > i0 {
				arrVal = append(arrVal, expression[i0:i])
			}
			arrCompute = append(arrCompute, '-')
			i0 = i + 1
		case '*':
			arrVal = append(arrVal, expression[i0:i])
			arrCompute = append(arrCompute, '*')
			i0 = i + 1
		}
	}
	if len(expression) > i0 {
		arrVal = append(arrVal, expression[i0:])
	}
	for i := range arrVal {
		v, ok := mapE[arrVal[i]]
		if ok {
			arrVal[i] = v
		}
	}
	var val int64 = 1
	if len(arrCompute) == len(arrVal) {
		if arrCompute[0] == '-' {
			val = -val
		}
		arrCompute = arrCompute[1:]
	}
	arr := []string{}
	for len(arrCompute) > 0 {
		switch arrCompute[0] {
		case '+', '-':
			vl, isNum := isNumber(arrVal[0])
			if isNum {
				val *= vl
				if len(arr) == 0 {
					valOnly += val
				} else {
					mapVal[getVal(arr)] += int(val)
					arr = []string{}
				}
			} else {
				arr = append(arr, arrVal[0])
				mapVal[getVal(arr)] += int(val)
				arr = []string{}
			}
			val = 1
			if arrCompute[0] == '-' {
				val = -1
			}
			arrVal = arrVal[1:]
			arrCompute = arrCompute[1:]
		default:
			vl, isNum := isNumber(arrVal[0])
			if isNum {
				val *= vl
			} else {
				arr = append(arr, arrVal[0])
			}
			arrVal = arrVal[1:]
			arrCompute = arrCompute[1:]
		}
	}
	vl, isNum := isNumber(arrVal[0])
	if isNum {
		val *= vl
		if len(arr) == 0 {
			valOnly += val
		} else {
			mapVal[getVal(arr)] += int(val)
		}
	} else {
		arr = append(arr, arrVal[0])
		mapVal[getVal(arr)] += int(val)
	}
	return mapVal, int(valOnly)
}
func getVal(arr []string) string {
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})
	return strings.Join(arr, "*")
}
func isNumber(s string) (int64, bool) {
	vl, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0, false
	}
	return vl, true
}