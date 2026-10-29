func partitionLabels(s string) []int {
    // Create an array to store the last occurrence of each character (a-z)
    lastOccurrence := make([]int, 26)
    for i, ch := range s {
        lastOccurrence[ch-'a'] = i
    }
    
    var result []int
    start, end := 0, 0
    
    for i, ch := range s {
        // Update the end of the current partition
        if lastOccurrence[ch-'a'] > end {
            end = lastOccurrence[ch-'a']
        }
        
        // If we've reached the end of the current partition
        if i == end {
            result = append(result, end-start+1)
            start = end + 1
        }
    }
    
    return result
}