class Solution {
    public String largestEven(String s) {
        int ind=-1;
        for(int i=s.length()-1;i>=0;i--){
            if(s.charAt(i)=='2'){
                ind=i;
                break;
            }
        }
        if(ind==-1) return "";
        return s.substring(0,ind+1);
        }
    }