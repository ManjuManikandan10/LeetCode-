type Work struct {
    difficulty int
    profit int
}

func maxProfitAssignment(difficulty []int, profit []int, worker []int) int {
    works := []Work{}
    for i := 0 ; i<len(profit) ; i++{
        currentDifficulty := difficulty[i]
        currentProfit := profit[i]
        works = append(works , Work{difficulty:currentDifficulty , profit:currentProfit})
    }

    sort.Slice(works, func(i, j int) bool {
        if works[i].difficulty == works[j].difficulty{
            return works[i].profit > works[j].profit
        }
        return works[i].difficulty < works[j].difficulty
    })


	maxProfit := make([]int, len(profit))
	maxProfit[0] = works[0].profit
	for i := 1; i < len(profit); i++ {
		maxProfit[i] = max(maxProfit[i-1], works[i].profit)
	}

    resp := 0
    for _, people := range worker {
        resp += binarySearch(works , people , maxProfit)
    }
    return resp
}

func binarySearch(task []Work , people int , maxProfit[]int) int {
    left , right := 0 , len(task)-1
    resp := 0
    for left <= right {
        mid := left + (right-left) / 2 

        if people >= task[mid].difficulty{
            resp = max(resp , maxProfit[mid])
            left = mid+1
        } else {
            right = mid -1
        }
    }
    return resp
}

func max(a,b int) int {
    if a > b {
        return a
    }
    return b
}