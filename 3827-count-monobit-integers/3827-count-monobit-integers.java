class Solution {
    public boolean isvalid(String str){
        if(str.length()==1){
            return true;
        }
        for(int i=0;i<str.length()-1;i++){
            if(str.charAt(i)!=str.charAt(i+1)){
                return false;
            }
        }
        return true;
    }
    public int countMonobit(int n) {
        int count=0;
        for(int i=0;i<=n;i++){
            if(isvalid(Integer.toBinaryString(i))){
                count++;
            }
        }
        return count;
    }
}