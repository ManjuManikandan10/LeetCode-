func validSubarraySize(nums []int, threshold int) int {


    ms := make([]int, 0)

    nextSmaller := make([]int, len(nums))
   
    for i := 0;i<len(nums);i++{
        for len(ms) > 0 && nums[ms[len(ms)-1]] > nums[i]{
            topIndex := ms[len(ms)-1]
            ms = ms[:len(ms)-1]

            nextSmaller[topIndex] = i - topIndex
        }

        ms = append(ms, i)
    }

    ms = make([]int, 0)

    previousSmaller := make([]int, len(nums))

    for i:= len(nums)-1; i >= 0;i--{
        for len(ms) > 0 && nums[ms[len(ms)-1]] > nums[i]{
            topIndex := ms[len(ms)-1]
            ms = ms[:len(ms)-1]

            previousSmaller[topIndex] = topIndex - i
        }

        ms = append(ms, i)
    }

    
    for i := range nums{
        toRight, toLeft := 0,0
        if nextSmaller[i] == 0{
            toRight = len(nums) - i - 1
        } else{
            toRight = nextSmaller[i] -1
        }

        if previousSmaller[i] == 0{
            toLeft = i
        } else{
            toLeft = previousSmaller[i]-1
        }
        k := toLeft + 1 + toRight
        area := k * nums[i]
        if area > threshold{
            return k
        } 
    }

    return -1

}