class Solution {
    
    //Classic Union-Find Algorithm to find Common Ancestor
    private int findParent(int[] parent, int index) {
        if(parent[index] == index) return index;
        return findParent(parent, parent[index]);
    }
   
    public boolean[] friendRequests(int n, int[][] restrictions, int[][] requests) {
        int resLength = restrictions.length, reqLength = requests.length;
        boolean[] result = new boolean[reqLength];
        int[] parent = new int[n];
        
        //Initially ith person's parent is i itself
        for(int i = 0; i < n; i++) parent[i] = i;
        

        for(int i = 0; i < reqLength; i++) {
            //finding the parents of the first person and second person of ith request
            int firstParent = findParent(parent, requests[i][0]);
            int secondParent = findParent(parent, requests[i][1]);
            
            //if they have same parents i.e. mutual friends they can be friends
            if(firstParent == secondParent) {
                result[i] = true;
                continue;
            }
            
            //iterating through the restrictions array to find whether the parents of first ans second person have a conflict 
            boolean flag = true;
            for(int j = 0; j < resLength; j++) {
                //finding parents of the restriction persons
                int firstRestriction = findParent(parent, restrictions[j][0]);
                int secondRestriction = findParent(parent, restrictions[j][1]);
                
                //if any of the parents are matching i.e. if the parents of first and second person have a mutual conflict they can't be friend
                if((firstRestriction == firstParent && secondRestriction == secondParent) || (secondRestriction == firstParent && firstRestriction == secondParent)) {
                    flag = false;
                    break;
                }
            }
            
            if(flag) {
                result[i] = true;
                parent[firstParent] = secondParent; //setting the common ancestor -> classic union find technique
            }
        }
        
        return result;
    }
}