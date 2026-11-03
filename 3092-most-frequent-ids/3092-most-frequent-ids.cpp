class Solution {
public:
    vector<long long> mostFrequentIDs(vector<int>& nums, vector<int>& freq) {
        //Please Upvote :)
        unordered_map<long long,long long>mp;
        int i,j;
        vector<long long>v(nums.size());
        priority_queue<pair<long long,long long>>pq;
        for(i=0;i<nums.size();i++){
            mp[nums[i]]+=freq[i];
            pq.push({mp[nums[i]],nums[i]});
            while(pq.size()>0 && mp[pq.top().second]!=pq.top().first){
                pq.pop();
            }
            v[i]=(pq.top().first);
        }
        return v;
    }
};