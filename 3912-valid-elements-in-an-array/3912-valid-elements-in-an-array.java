class Solution {
    public List<Integer> findValidElements(int[] nums) {
        List<Integer> list= new ArrayList<>();
        list.add(nums[0]);
        boolean numleft=true;
        boolean numright=true;
        for(int i=1;i<nums.length-1;i++){
            for(int j=0;j<i;j++){
                if(nums[j]>=nums[i])
                    numleft=false;
            }
            for(int k=i+1;k<nums.length;k++){
                if(nums[k]>=nums[i])
                    numright=false;
            }
            if(numleft || numright){
                list.add(nums[i]);
            }
            numleft=true;
            numright=true;
        }
        if(nums.length>1)
            list.add(nums[nums.length-1]);
        return list;
    }
}