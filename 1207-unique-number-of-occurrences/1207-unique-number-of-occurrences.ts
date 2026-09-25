function uniqueOccurrences(arr: number[]): boolean {
    // declare frequency map
    const freq = new Map<number, number>();
    // calculate frequency map of array numbers
    arr.forEach((num) => freq.set(num, freq.has(num) ? freq.get(num) + 1 : 1));
    // declare set
    const set = new Set<number>();
    // iterate over frequency map
    for(const [key, value] of freq) {
        // if we found not unique value we return false
        if(set.has(value)) return false;
        // if not add value to set
        set.add(value)
    }
    // return true if all frequencies unique
    return true;
};