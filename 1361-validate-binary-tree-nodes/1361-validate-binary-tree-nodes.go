func validateBinaryTreeNodes(n int, leftChild []int, rightChild []int) bool {
    parent := make([]int, n)
    for i:=0; i<n; i++ {
        parent[i] = -1
    }

    for i:=0; i<n; i++ {
        if leftChild[i] != -1 {
            if parent[leftChild[i]] != -1 {
                return false
            }
            parent[leftChild[i]] = i
        }

        if rightChild[i] != -1 {
            if parent[rightChild[i]] != -1 {
                return false
            }
            parent[rightChild[i]] = i
        }
    }

    count:=0
    for i:=0; i<n; i++ {
        if parent[i] == -1 {
            count++
        }
    }

    if count == 1 {
        return !checkCyclic(parent, n)
    }

    return false
}

func checkCyclic(p []int, n int) bool {
    var vstd []bool
    for i:=0; i<n; i++ {
        vstd=make([]bool, n)
        j:=i
        for p[j]!=-1 && vstd[j]==false {
            vstd[j]=true
            j=p[j]
            
        }

        if vstd[j] {
            return true
        }
    }

    return false
}