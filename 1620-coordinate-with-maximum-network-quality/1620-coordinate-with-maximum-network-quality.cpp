class Solution {
public:
    vector<int> bestCoordinate(vector<vector<int>>& towers, int radius) {
        // Calculate the range
        int minx =
            (*std::min_element(towers.begin(), towers.end(),
                             [](auto& a, auto& b) { return a[0] < b[0]; }))[0];
        int maxx =
            (*std::min_element(towers.begin(), towers.end(),
                             [](auto& a, auto& b) { return a[0] > b[0]; }))[0];
        int miny =
            (*std::min_element(towers.begin(), towers.end(),
                             [](auto& a, auto& b) { return a[1] < b[1]; }))[1];
        int maxy =
            (*std::min_element(towers.begin(), towers.end(),
                             [](auto& a, auto& b) { return a[1] > b[1]; }))[1];

        vector<int> prevMax{0, 0, 0};
        // Go through all options and compare
        for (int x{minx}; x <= maxx; x++) {
            for (int y{miny}; y <= maxy; y++) {
                int cur{0};
                for (auto& t : towers) {
                    // continue distance
                    double d = std::sqrt((x - t[0]) * (x - t[0]) +  (y - t[1]) * (y - t[1]));
                    if (d > radius) continue;
                    cur += (t[2] / (1 + d));
                }
                if (cur > prevMax[2]) {
                    prevMax = {x, y, cur};
                }
            }
        }
            
        return {prevMax[0], prevMax[1]};                
    }
};