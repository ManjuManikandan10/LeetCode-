def find_kth_number(m, n, k)
    ans = (1..k).bsearch { |i| count(m,n,i) >= k }
end

def count(m,n,i)
    total = 0
    (1..m).each { |j| total += [i/j, n].min }
    total
end