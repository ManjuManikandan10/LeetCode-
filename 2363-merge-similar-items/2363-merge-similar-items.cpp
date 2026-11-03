class Solution {
public:
    vector<vector<int>> mergeSimilarItems(vector<vector<int>>& items1, vector<vector<int>>& items2) {

        vector<vector<int>> result;
        map<int, int> mpp;

        for(int i = 0; i < items1.size(); i++){
            mpp[items1[i][0]] += items1[i][1];
        }

        for(int j = 0; j < items2.size(); j++){
            mpp[items2[j][0]] += items2[j][1];
        }

        for(auto item : mpp){
            result.push_back({item.first, item.second});
        }

        return result;
    }
};