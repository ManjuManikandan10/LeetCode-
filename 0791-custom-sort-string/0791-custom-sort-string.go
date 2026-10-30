func customSortString(order string, s string) string {
    mp := make(map[rune]int)
    for _, r := range s {
        mp[r]++
    }
    var b strings.Builder
    for _, r := range order {
        b.WriteString(strings.Repeat(string(r), mp[r]))
        mp[r] = 0
    }
    for k, v := range mp {
        if v != 0 {
            b.WriteString(strings.Repeat(string(k), v))
        }
    }
    return b.String()
}