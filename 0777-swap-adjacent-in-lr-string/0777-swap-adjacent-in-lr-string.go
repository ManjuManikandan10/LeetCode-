
func canTransform(start string, result string) bool {
	if len(start) != len(result) {
		return false
	}
	is := 0
	ir := 0
	for is < len(start) && ir < len(result) {
		if start[is] == 'X' {
			is++
			continue
		}
		if result[ir] == 'X' {
			ir++
			continue
		}
		if start[is] != result[ir] {
			return false
		}
		is++
		ir++
	}
	for i := ir; i < len(result); i++ {
		if result[i] != 'X' {
			return false
		}
	}
	for i := is; i < len(start); i++ {
		if start[i] != 'X' {
			return false
		}
	}
	countS := 0
	countR := 0
	for i := 0; i < len(start); i++ {
        if start[i] == 'X' {
			countS++
		}
		if result[i] == 'X' {
			countR++
		}
		if (result[i] == 'L'||start[i] == 'L') && countR > countS {
			return false
		}
		if (result[i] == 'R' ||start[i] == 'R') && countR < countS {
			return false
		}
	}
	return true
}