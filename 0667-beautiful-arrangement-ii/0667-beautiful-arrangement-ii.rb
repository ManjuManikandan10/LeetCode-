# @param {Integer} n
# @param {Integer} k
# @return {Integer[]}
def construct_array(n, k)
  Enumerator.produce([k / 2 + 1, k % 2 * 2 - 1]) { [_1 + _2, (0 <=> _2) - _2] }.take(k + 1).map(&:first) + ((k + 2)..n).to_a
end