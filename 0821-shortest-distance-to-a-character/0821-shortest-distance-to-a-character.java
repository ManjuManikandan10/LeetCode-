class Solution {
    public int[] shortestToChar(String s, char c) {
        int n = s.length();
        int[] result = new int[n];

        int prevC = -1;   
        int currC = -1;  

        int scanIndex = 0; 
        int fillIndex = 0; 

        while (fillIndex < n) {

            while (scanIndex < n) {
                if (s.charAt(scanIndex) == c) {
                    prevC = currC;
                    currC = scanIndex;
                    break;
                }
                scanIndex++;
            }

            while (fillIndex < n) {
                int distLeft  = (prevC != -1) ? fillIndex - prevC : n;
                int distRight = Math.abs(currC - fillIndex);

                result[fillIndex] = Math.min(distLeft, distRight);

                if (result[fillIndex] == 0) break;
                fillIndex++;
            }

            scanIndex++;
            fillIndex++;
        }

        return result;
    }
}