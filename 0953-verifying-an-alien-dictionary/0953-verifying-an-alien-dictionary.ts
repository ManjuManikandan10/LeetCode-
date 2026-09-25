const ASCII_OFFSET = 'a'.charCodeAt(0)
const ASCII_LAST = 'z'.charCodeAt(0)

const char2index = (char: string) => char.charCodeAt(0) - ASCII_OFFSET
const getCharOrder = (char: string, dict: number[]) => dict[char2index(char)]

function isAlienSorted(words: string[], order: string): boolean {
  const orderSeq: number[] = Array(ASCII_LAST - ASCII_OFFSET).fill(0)
  for (let i = 0; i < order.length; i++) {
    const j = char2index(order[i])
    orderSeq[j] = i
  }

  for (let i = 0; i < words.length - 1; i++) {
    const w1 = words[i]
    const w2 = words[i+1]
    const minLen = Math.min(w1.length, w2.length)
    for (let j = 0; j < minLen; j++) {
      const c1 = w1[j]
      const c2 = w2[j]
      if (c1 === c2) continue
      
      const o1 = getCharOrder(c1, orderSeq)
      const o2 = getCharOrder(c2, orderSeq)
      if (o1 < o2) {
        break
      }

      return false
    }

    if (w1.length > w2.length && w1.slice(0, minLen) === w2.slice(0, minLen)) {
      // edge case
      return false
    }
  }

  return true
};