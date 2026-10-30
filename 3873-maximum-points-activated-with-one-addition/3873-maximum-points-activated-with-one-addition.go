import "strconv"
type Point struct {
    x string 
    y string 
    key string
}

func NewPoint(x string, y string) *Point{
    
    return &Point{
        x:x, 
        y:y,
        key:x + ";" + y,
    }
}

func (p *Point) Key() string {
    return p.key 
}

func preprocess(points[][]int) (map[string][]*Point, []*Point) {
    adj := map[string][]*Point{}
    ps := []*Point{}
    for _, point := range points {
        x_key := NewPoint(strconv.Itoa(point[0]), "y").Key()
        y_key := NewPoint("x", strconv.Itoa(point[1])).Key()
        p := NewPoint(strconv.Itoa(point[0]), strconv.Itoa(point[1]))
        //fmt.Println(point, p.x, p.y, p.Key())
        adj[x_key] = append(adj[x_key], p)
        adj[y_key] = append(adj[y_key], p)
        ps = append(ps, p)
    }
    return adj, ps
}



func maxActivated(points [][]int) int {
    adj, ps := preprocess(points)
    seen := map[string]bool{}
    first, second := 0, 0 

    var traverse func(point *Point) int
    traverse = func(point *Point) int {
        if seen[point.Key()] {
            return 0 
        }
        seen[point.Key()] = true 
        if point.x == "x" || point.y == "y"{
            val := 0 
            for _, next := range adj[point.Key()] {
                if !seen[next.Key()] {
                    val += traverse(next)
                }
            }
            return val 
        } else {
            val := 1 
            for _, next := range []*Point{NewPoint(point.x, "y"), NewPoint("x", point.y)} {
                if !seen[next.Key()] {
                    val += traverse(next)
                }
            }
            return val
        }
    }
    
    for _, point := range ps {
        val := traverse(point)
        if val > first {
            first, second = val, first 
        } else if val > second {
            second = val 
        }
    }
    return first + second + 1
}