func findReplaceString(s string, indices []int, sources []string, targets []string) string {
    indexes := make(map[int][]int, len(indices))
    for i,v := range indices {
        indexes[v] = append(indexes[v], i)
    }
    
    result := make([]string, 0, len(s))
    i := 0
    for i < len(s) {
        current := string(s[i])
        parr, ok := indexes[i]
        
        wasReplaced := false
        if ok {
            for _, p := range parr {
                source := sources[p]
                target := targets[p]
                sl := len(source)
                if i+sl <= len(s) && s[i:i+sl] == source{
                    result = append(result, target)
                    i += sl
                    wasReplaced = true
                    break
                }
            }
            
        }
        
        if !wasReplaced {
            result = append(result, current)
            i++
        } 
    }
 
    return strings.Join(result, "")
}