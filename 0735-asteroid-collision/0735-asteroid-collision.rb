def asteroid_collision(asteroids)
  stack = []

  asteroids.each do |asteroid|
    if asteroid > 0
      stack << asteroid
    else
      while !stack.empty? && stack.last > 0 && stack.last < asteroid.abs
        stack.pop
      end

      if stack.empty? || stack.last < 0
        stack << asteroid
      elsif stack.last == asteroid.abs
        stack.pop
      end
    end
  end

  stack
end