class Solution {
    public List<List<Integer>> largeGroupPositions(String s) 
    {
        List<List<Integer>> result = new ArrayList<>();
        int front = 0;
        int rear = s.length()-1;
        int stIndex = 0;
        int lstIndex = 0;
        while(front < rear)
        {
            stIndex = front;
            int count = 1;
            while(front < rear && s.charAt(front) == s.charAt(front+1))
            {
                count++;
                front++;
            }
            if(count >= 3)
            {
                List<Integer> list = new ArrayList<>();
                list.add(stIndex);
                list.add(front);
                result.add(list);
            }
            
            front++;
        }
        return result;
    }
}