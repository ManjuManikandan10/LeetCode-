class Solution {
    public void duplicateZeros(int[] arr) {
        int n = arr.length;
        int[] temp = new int[n];
        int j = 0;

        for(int i = 0; i < n && j < n; i++){
            temp[j++] = arr[i];

            if(arr[i] == 0 && j < n){
                temp[j++] = 0;
            }
        }

        for(int i = 0; i < n; i++){
            arr[i] = temp[i];
        }
    }
}