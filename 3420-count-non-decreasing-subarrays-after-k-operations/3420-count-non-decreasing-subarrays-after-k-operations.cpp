class Solution {
public:
    long long countNonDecreasingSubarrays(vector<int>& nums, int k) {
        int n = nums.size();

        vector<int64_t> pre(n + 1);
        for (int i = 0; i < n; i++) {
            pre[i + 1] = pre[i] + nums[i];
        }

        vector<int> r(n + 1, n);
        vector<int> stk;

        for (int i = n - 1; i >= 0; i--) {
            while (stk.size() && nums[i] >= nums[stk.back()]) {
                stk.pop_back();
            }
            r[i] = stk.size() ? stk.back() : n;
            stk.push_back(i);
        }

        vector<vector<int>> up(20, vector<int>(n + 1, n));
        vector<vector<int64_t>> cost(20, vector<int64_t>(n + 1));

        for (int i = 0; i < n; i++) {
            up[0][i] = r[i];

            int len = r[i] - i;
            int64_t s = pre[r[i]] - pre[i];

            cost[0][i] = 1LL * len * nums[i] - s;
        }

        for (int i = 1; i < 20; i++) {
            for (int j = 0; j < n; j++) {
                up[i][j] = up[i - 1][up[i - 1][j]];
                cost[i][j] = cost[i - 1][j] + cost[i - 1][up[i - 1][j]];
            }
        }

        auto get = [&](int i) -> int {
            int64_t K = k;

            for (int x = 19; x >= 0; x--) {
                if (cost[x][i] <= K) {
                    K -= cost[x][i];
                    i = up[x][i];
                }
            }

            if (i == n) return n - 1;

            int l = i + 1, h = r[i] - 1;

            while (l <= h) {
                int mid = (l + h) >> 1;

                int64_t s = pre[mid + 1] - pre[i + 1];
                int len = mid - i;

                int64_t nrem = K - (1LL * len * nums[i] - s);

                if (nrem >= 0) {
                    l = mid + 1;
                } else {
                    h = mid - 1;
                }
            }

            return l - 1;
        };

        int64_t ans = 0;

        for (int i = 0; i < n; i++) {
            ans += get(i) - i + 1;
        }

        return ans;
    }
};