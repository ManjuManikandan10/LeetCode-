func isIdealPermutation(nums []int) bool {
    m := make(map[int]int)
    for i:=0;i<len(nums);i++ {
        m[nums[i]] = i
    }
    for i:=0;i<len(nums);i++ {
        if m[i] > i+1 || m[i] < i-1{
            return false
        }
    }
    return true
}