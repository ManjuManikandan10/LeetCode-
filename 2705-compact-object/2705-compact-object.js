/**
 * @param {Object|Array} obj
 * @return {Object|Array}
 */
var compactObject = function (obj) {
    function dfs(obj) {
        if (!obj) return false
        if (typeof obj !== 'object') return obj

        if (Array.isArray(obj)) {
            const newArr = []
            for (let item of obj) {
                const subRes = dfs(item)

                if (subRes) {
                    newArr.push(subRes)
                }

            }
            return newArr
        }

        const newObj = {};

        for (let key in obj) {
            const subRes = dfs(obj[key])

            if (subRes) {
                newObj[key] = subRes
            }
        }
        return newObj
    }
    return dfs(obj)
};