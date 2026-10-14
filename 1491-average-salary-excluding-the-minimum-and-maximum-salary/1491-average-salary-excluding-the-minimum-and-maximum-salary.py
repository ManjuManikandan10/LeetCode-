class Solution:
    def average(self, salary: List[int]) -> float:
        minsal=min(salary)
        maxsal=max(salary)
        salary.pop(salary.index(minsal))
        salary.pop(salary.index(maxsal))
        sum1=sum(salary)
        average=sum1/len(salary)
        return float(average)
        #please do upvote it will encourage me alot.