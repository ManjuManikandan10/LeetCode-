class Solution {
    const int inf = 1e9;

    int cost_to_palindrome(string s) {
        int n = s.size();
        int cost = 0;

        for(int i=0, j = n-1;i<j;i++, j--) {
            if(s[i] != s[j]) {
                cost++;
            }
        }

        return cost;
    }

public:
    int minimumChanges(string s, int k) {
        int n = s.size();
        
        vector<vector<int>>cost(n, vector<int>(n, inf));
        // cost[i][j] represents the minimum cost to make the substring [i:j] semi-palindrome

        for(int i=0;i<n;i++) {
            for(int j=i+1;j<n;j++) {
                int l = j-i+1;
                for(int d=1;d<l;d++) {
                    if(l%d != 0) continue;

                    int cur_cost = 0;
                    for(int st=i;st<i+d;st++) {
                        string cur;
                        for(int k=st;k<=j;k+=d) {
                            cur.push_back(s[k]);
                        }
                        
                        cur_cost += cost_to_palindrome(cur);
                    }

                    cost[i][j] = min(cost[i][j], cur_cost);
                }
            }
        }

        vector<vector<int>>dp(n, vector<int>(k+1, inf));
        // dp[i][j] represents the minimum cost to partition the prefix [0:i] into j parts

        for(int i=0;i<n;i++) {
            dp[i][1] = cost[0][i];
            for(int j=2;j<=k;j++) {
                for(int last=0;last<i;last++) {
                    dp[i][j] = min(dp[i][j], dp[last][j-1]+cost[last+1][i]);
                }
            }
        }

        return dp[n-1][k];
    }
};