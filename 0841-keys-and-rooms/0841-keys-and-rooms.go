func canVisitAllRooms(rooms [][]int) bool {
    isOpen := make([]bool, len(rooms))
    
    var open func(key int) 
    open = func(key int) {
        if isOpen[key] == true {
            return
        }
        isOpen[key] = true
        for _, k := range rooms[key] {
            open(k)
        }
    }
    open(0)
    for _, op := range isOpen {
        if !op {
            return false
        }
    }
    return true
}