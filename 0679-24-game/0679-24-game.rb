def judge_point24(a) =

  !a[1] ? a==[24] : '+-*/'.chars.any?{|o| a.permutation.any?{
    judge_point24 _1 << _1.pop(2).map(&:to_r).reduce(o) rescue nil } }