func consecutiveNumbersSum(n int) (result int) {
    for x, k := float64(n), 1.0; x > 0; x = (((2 * float64(n)) / k) - k + 1.0) / 2.0 {
        if x == math.Ceil(x) {
            result++
        }
        k++
    }
    return 
}