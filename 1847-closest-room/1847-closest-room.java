class Solution {
    public int[] closestRoom(int[][] rooms, int[][] queries) {
        
        int numRoom = rooms.length;
        int numQuery = queries.length;
        
        
        for(int i = 0; i < numQuery; i++){
			//append index
            queries[i] = new int[]{queries[i][0], queries[i][1], i};
        }
        
        //sort the roos and queries in increasing order of room size
        Arrays.sort(rooms, (a, b) -> (a[1] != b[1] ? (a[1] - b[1]) : (a[0] - b[0])));
        Arrays.sort(queries, (a, b) -> (a[1] != b[1] ? (a[1] - b[1]) : (a[0] - b[0])));
        
        TreeSet<Integer> roomIds = new TreeSet();
        int[] result = new int[numQuery];
        
        int j = numRoom - 1;
        for(int i = numQuery - 1; i >= 0; i--){
            
            int currRoomId = queries[i][0];
            int currRoomSize = queries[i][1];
            int currQueryIndex = queries[i][2];
            
            //add all the which size is greater or equal than query room size
            while(j >= 0 && rooms[j][1] >= currRoomSize){
                roomIds.add(rooms[j--][0]);
            }
            
            //found room with sufficient size,
            if(roomIds.contains(currRoomId)){
                //directly add this room
                result[currQueryIndex] = currRoomId;
                continue;
            }
            
            //next room which is just greater that current roomId
            Integer nextRoomId = roomIds.higher(currRoomId);
            
            //previous roomId which is just smaller than current roomId
            Integer prevRoomId = roomIds.lower(currRoomId);
            
            if(nextRoomId == null && prevRoomId == null){
                //no nearest roomId 
                result[currQueryIndex] = -1;
                
            }else if(nextRoomId == null){
                
                result[currQueryIndex] = prevRoomId;
                
            }else if(prevRoomId == null){
                
                result[currQueryIndex] = nextRoomId;
                
            }else if(nextRoomId != null && prevRoomId != null){
                //get the nearest roomId
                if((currRoomId - prevRoomId) <= (nextRoomId - currRoomId)){
                    result[currQueryIndex] = prevRoomId;
                }else{
                    result[currQueryIndex] = nextRoomId;
                }
            }
            
        }
        
        return result;
    }
}