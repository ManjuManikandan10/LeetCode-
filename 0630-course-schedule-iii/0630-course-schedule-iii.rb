# @param {Integer[][]} courses
# @return {Integer}
def schedule_course(courses)
    pq = PriorityQueue.new
    courses.sort! {|a, b| a[1] <=> b[1]}
    t = 0
    for course in courses do
        pq.push(course, course[0])
        t += course[0]
        if t>course[1] 
            popped = pq.pop
            t-=popped[0]
        end
    end
    pq.size
end