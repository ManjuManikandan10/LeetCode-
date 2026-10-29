func numRabbits(answers []int) int {
    freq := make(map[int]int) 
    for _, k := range answers {
        freq[k+1]++
    }
    
    var c int
    for k, v  := range freq {
        if k >= v {
            c += k
        } else {
            c += v - v % k
            if v % k > 0 {
                c += k
            }
        }
    }
    
    return c
}