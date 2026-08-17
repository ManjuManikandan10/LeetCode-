class Solution {
    public int nearestDrone(int[][] drones, int[] target) {
        int ans = -1;
        int minDist = Integer.MAX_VALUE;
        int i = 0;

        for(int[] drone :  drones) {
            int dist = Math.abs(drone[0] - target[0]) + Math.abs(drone[1] - target[1]);

            if(dist <= drone[2] && minDist > dist) {
                ans = i;
                minDist = dist;
            }

            i++;
        }

        return ans;
    }
} 