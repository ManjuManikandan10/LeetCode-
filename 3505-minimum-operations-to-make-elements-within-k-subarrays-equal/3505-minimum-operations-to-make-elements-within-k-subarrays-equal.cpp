class Solution {
public:
    vector<vector<long long>>dp;
    long long solve(int i, vector<long long>&ops, int x, int k, int n){
        if(k <= 0) return 0;
        if(i >= n){
            if(k > 0) return LLONG_MAX;
            else return 0;
        }

        if(dp[i][k] != -1) return dp[i][k];

        // take current subarray cost
        long long min_cost = solve(i+x, ops, x, k-1, n);
        if(min_cost != LLONG_MAX) min_cost += ops[i] ;
        //dont take current subarray cost
        min_cost = min(min_cost, solve(i+1, ops, x, k, n));

        return dp[i][k] = min_cost ;

    }
    long long minOperations(vector<int>& nums, int x, int k) {
        int n = nums.size();
        int mx_sub = (n/x);
        if(mx_sub < k) return -1;

        vector<long long>ops;

        // main problem is that we can't do operations calculation for current sliding window becaue it would cause tle

        // to deal with this we will have to use standard method of finding median in o(1) whihch is to use - 2 priority queues or 2 multisets, because in this method while adding elements we can calculate operations simultaneously

        multiset<int>lower, higher ;
        long long sumLow = 0, sumHigh =0 ;

        // HELPING FUNCTIONS FOR MEDIAN
        auto rebalance = [&](){
            int need = (lower.size() + higher.size() +1)/2 ;
            while(lower.size() > need){
                auto it = prev(lower.end());
                int val = *it;
                lower.erase(it);
                higher.insert(val);
                sumLow -= val;
                sumHigh += val;
            }

            while(lower.size() < need){
                int val = *higher.begin();
                lower.insert(val);
                higher.erase(higher.begin());
                sumLow += val;
                sumHigh -= val;
            }
        };

        auto add = [&](int val){
            if(!lower.empty() && val <= *lower.rbegin()){
                lower.insert(val);
                sumLow += val;
            }
            else{
                higher.insert(val);
                sumHigh += val;
            }
            rebalance();
        };

        // To remove elements from our current sets so new sliding window median can be calculated

        auto remove = [&](int val){
            auto it = lower.find(val);

            if(it != lower.end()){
                lower.erase(it);
                sumLow -= val;
            }
            else{
                higher.erase(higher.find(val));
                sumHigh -= val;
            }
            rebalance();
        };

        // IMPORTANT FUNCTION TO CALCULATE COST FOR CURRENT MEDIAN
        auto getcost = [&]() -> long long{
            long long median = *lower.rbegin();
            return median*lower.size()- sumLow + sumHigh - higher.size()*median ;
        };
    
        for (int i=0; i<x; i++) add(nums[i]);
        ops.push_back(getcost()); // cost of 1st window
        for(int i= x; i<n; i++){
            remove(nums[i-x]);
            add(nums[i]);
            ops.push_back(getcost());
        }
        dp.resize(ops.size(), vector<long long>(k+1,-1));
        int sz = ops.size();

        long long ans = solve(0, ops, x, k, sz);
        return ans;
    }
};
