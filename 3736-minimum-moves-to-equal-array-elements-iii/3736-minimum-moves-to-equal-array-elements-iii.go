func minMoves(nums []int) int {
    m := slices.Max(nums)
    moves := 0
    for _, n := range nums {
        moves += (m - n)
    }
    return moves
}