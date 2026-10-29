func reorganizeString(s string) string {
    hashMap := make(map[rune]int)

    // Count frequencies
    for _, val := range s {
        hashMap[val]++
    }

    // Find the character with the maximum frequency
    maxCount := 0;
    maxChar := ' '

    for i, val := range hashMap {
        
        if val > maxCount {
            maxCount = val
            maxChar = i
        }
    }

    // Check if reorganization is possible
    if maxCount > (len(s)+1)/2 {
        return ""
    }

    result := make([]byte, len(s))
    index := 0

    // Place the highest frequency character at even positions
    for hashMap[maxChar] > 0 {
        result[index] = byte(maxChar)
        index += 2
        hashMap[maxChar]--
    }

    // Fill other characters 
    for i := range hashMap {
        for hashMap[i] > 0 {

            if index >= len(result){
                index = 1 
            }

            result[index] = byte(i)
            hashMap[i]--
            index+=2
        }
    }

     return string(result)
}