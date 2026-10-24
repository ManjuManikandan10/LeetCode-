def count_of_atoms(formula) =

  (?(+formula).scan(/(\W|[A-Z][a-z]*)(\d*)/).reduce([]){|s,(t,n)|
    u = proc{ s.last[_1] += [n.to_i,1].max*_2 }
    t=='(' ? s << Hash.new(0) : t==')' ? s.pop.each(&u) : u[t,1]
    s }.pop.sort.each{_1.pop if _1[1]<2}.join