
func preimageSizeFZF(k int) int {
	mn := 0
	mx := math.MaxInt64
	vm := getSizeFZF(mx)
	if vm == k {
		return 5
	}
	if vm < k {
		return 0
	}
	mid := (mn + mx) / 2
	vm = getSizeFZF(mid)
	for vm != k {
		if vm > k {
			mx = mid
		} else {
			mn = mid
		}
		if mid+1 == mx || mid-1 == mn {
			return 0
		}
		mid = (mn + mx) / 2
		vm = getSizeFZF(mid)
	}

	return 5
}
func getSizeFZF(n int) int {
	rs := 0
	n5 := n
	for n5 >= 5 {
		rs += n5 / 5
		n5 /= 5
	}
	return rs
}