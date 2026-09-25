function numEquivDominoPairs(dominoes: number[][]): number {
    let count = 0
    let myMap = new Map()
    for(let i = 0; i < dominoes.length; i++){
        let [a,b] = dominoes[i]
        let key = a < b ? `${a}${b}` : `${b}${a}`
        if(myMap.has(key)){
            count += myMap.get(key)
        }
        myMap.set(key, myMap.get(key) + 1 || 1)
    }
    return count
};