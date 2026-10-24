def get_importance(employees, id)
  ans = 0
  manager = employees.find { |emp| emp.id == id }
  queue = [manager.subordinates]
  ans = manager.importance
  while !queue.empty?
    curr = queue.shift
    curr.each do |id|
      emp = employees.find { |e| e.id == id }
      ans += emp.importance
      queue.push(emp.subordinates) unless emp.subordinates.empty?
    end
  end
  ans
end