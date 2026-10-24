def maximum_swap(num)
  num_str = num.to_s.chars
  last = Array.new(10, -1)
  
  num_str.each_with_index do |digit, i|
    last[digit.to_i] = i
  end

  num_str.each_with_index do |digit, i|
    (9).downto(digit.to_i + 1) do |d|
      if last[d] > i
        num_str[i], num_str[last[d]] = num_str[last[d]], num_str[i]
        return num_str.join.to_i
      end
    end
  end

  num
end