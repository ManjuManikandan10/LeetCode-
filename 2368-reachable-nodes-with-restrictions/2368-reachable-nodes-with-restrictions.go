func reachableNodes(n int, edges [][]int, restricted []int) int {
    adjList := make([][]int, n)
    for _, v := range edges {
        adjList[v[0]] = append(adjList[v[0]], v[1])
        adjList[v[1]] = append(adjList[v[1]], v[0])
    }
    
    seen := make(map[int]bool)
    seen[0] = true
    for _, v := range restricted {
        seen[v] = true
    }
    stack := make([]int, 0)
    stack = append(stack, adjList[0]...)
    
    ans := 1
    
    var popLeft int
    for len(stack) > 0 {
        stack, popLeft = stack[1:], stack[0]
        if seen[popLeft] {
            continue
        }
        seen[popLeft] = true
        ans+=1
        stack = append(stack, adjList[popLeft]...)
    }
    return ans
}