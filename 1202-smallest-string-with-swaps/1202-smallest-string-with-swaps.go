import "sort"

func smallestStringWithSwaps(s string, pairs [][]int) string {
    p := make([]int, len(s))
    for i := 0; i < len(s); i++ {
        p[i] = i
    }
    
    var find func(x int) int
    find = func(x int) int {
        if p[x] != x {
            p[x] = find(p[x])
        }
        return p[x]
    }
    
    var union func(x, y int)
    union = func(x, y int) {
        px := find(x)
        py := find(y)
        p[py] = px
    }
    
    for _, pair := range pairs {
        union(pair[0], pair[1])
    }
    
    dict := make(map[int][]byte)
    for i, v := range p {
        pi := find(v)
        dict[pi] = append(dict[pi], s[i])
    }
    
    for i := range dict {
        sort.Slice(dict[i], func(a int, b int) bool {
            return dict[i][a] < dict[i][b]
        })
    }
    
    res := make([]byte, 0, len(s))
    for _, v := range p {
        pi := find(v)
        res = append(res, dict[pi][0])
        dict[pi] = dict[pi][1:]
    }
    
    return string(res)
}