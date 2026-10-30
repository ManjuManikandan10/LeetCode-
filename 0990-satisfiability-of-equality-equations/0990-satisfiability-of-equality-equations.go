type UnionFind struct {
    Parents []int
    Ranks   []int
}

func newUF(size int) *UnionFind {
    parents := make([]int, 0)
    for n := range size {
        parents = append(parents, n)
    }

    ranks := make([]int, size)

    return &UnionFind{
        Parents: parents,
        Ranks: ranks,
    }
}

func (uf *UnionFind) Find(x int) int {
    if uf.Parents[x] != x {
        uf.Parents[x] = uf.Find(uf.Parents[x])
    }
    return uf.Parents[x]
}

func (uf *UnionFind) Union(x, y int) {
    xRoot := uf.Find(x)
    yRoot := uf.Find(y)

    if uf.Ranks[xRoot] > uf.Ranks[yRoot] {
        uf.Parents[yRoot] = xRoot
    } else {
        uf.Parents[xRoot] = yRoot
        if uf.Ranks[xRoot] == uf.Ranks[yRoot] {
            uf.Ranks[yRoot]++
        }
    }
}

func equationsPossible(equations []string) bool {
    uf := newUF(26) // cuz its lowercase, it's constrained to 26.

    for _, eq := range equations {
        u, size := utf8.DecodeRuneInString(eq)

	    eq = eq[size:]
	    rune2, size2 := utf8.DecodeRuneInString(eq)

	    eq = eq[size2:]
	    rune3, _ := utf8.DecodeRuneInString(eq)
        
        v, _ := utf8.DecodeLastRuneInString(eq)

        if rune2 == '=' && rune3 == '=' {
            uInt := int(u - 'a')
            vInt := int(v - 'a')

        
            uf.Union(uInt, vInt)
        }
    }


    for _, eq := range equations {
        u, size := utf8.DecodeRuneInString(eq)

	    eq = eq[size:]
	    rune2, size2 := utf8.DecodeRuneInString(eq)

	    eq = eq[size2:]
	    rune3, _ := utf8.DecodeRuneInString(eq)
        
        v, _ := utf8.DecodeLastRuneInString(eq)

        if rune2 == '!' && rune3 == '=' {
            uInt := int(u - 'a')
            vInt := int(v - 'a')

        
            if uf.Find(uInt) == uf.Find(vInt) { // it means there is a connection but should be
                return false
            }
        }
    }

    return true
}