class Solution {
    int[][] arr;
    int n,m;

    int[] dirx = {1, 0, -1, 0};
    int[] diry = {0, -1, 0, 1};

    private boolean valid(int x, int y){
        return x >=0 && x<n && y>=0 && y<m;
    }
    public int containVirus(int[][] isInfected) {
        arr = isInfected;
        n = arr.length;
        m = arr[0].length;

        int ans = 0;

        boolean[][] vis = new boolean[n][m];

        while(true){
            int[] curr_comp = next_comp();

            if(curr_comp[0] == -1 && curr_comp[1] == -1) break;

            int temp = dfs2(curr_comp[0], curr_comp[1]);
            ans += temp;
            

            for(int i=0 ; i<n ; i++){
                for(int j=0 ; j<m ; j++){
                    if(arr[i][j] == 1 && !vis[i][j]){
                        dfs3(i, j, vis);
                    }
                }
            }

            for(boolean[] row : vis) Arrays.fill(row, false);
            
        }

        return ans;
    }

    private void dfs3(int i, int j, boolean[][] vis){
        vis[i][j] = true;

        for(int k=0 ; k<4 ;k++){
            int newi = i+ dirx[k];
            int newj  = j+ diry[k];

            if(valid(newi, newj) && !vis[newi][newj]){
                if(arr[newi][newj] == 1) dfs3(newi, newj, vis);
                else if(arr[newi][newj] == 0){
                    vis[newi][newj] = true;
                    arr[newi][newj] = 1;
                }
            }

        }
    }

    private int dfs2(int i, int j){
        arr[i][j] = -1;

        int ans = 0;

        for(int k=0 ; k<4 ;k++){
            int newi = i+ dirx[k];
            int newj  = j+ diry[k];

            if(valid(newi, newj)){
                if(arr[newi][newj]== 1)  ans += dfs2(newi, newj);
                else if(arr[newi][newj] == 0) ans++;
            }
        }

        return ans;
    }

    private int dfs(int i, int j, Queue<int[]> q, boolean[][] vis){
        vis[i][j] = true;
        q.add(new int[]{i, j});

        int ans = 1;

        for(int k=0 ; k<4 ;k++){
            int newi = i + dirx[k];
            int newj = j + diry[k];

            if(valid(newi, newj) && !vis[newi][newj] && arr[newi][newj] == 1){
                ans += dfs(newi, newj, q, vis);
            }
        }

        return ans;
    }

    private int[] next_comp(){
        int size = 0,ci=-1, cj=-1;
        boolean[][] vis = new boolean[n][m];
        Queue<int[]> q = new LinkedList<>();

        for(int i=0 ; i<n ; i++){
            for(int j=0 ; j<m; j++){
                if(arr[i][j] == 1 && !vis[i][j]){
                    int component_size = dfs(i, j, q, vis);

                    int curr_size = 0;

                    boolean[][] vis2 = new boolean[n][m];

                    while(component_size-- > 0){
                        int[] curr = q.poll();
                        int x = curr[0], y = curr[1];

                        for(int k=0 ; k<4 ; k++){
                            int newi = x + dirx[k];
                            int newj = y + diry[k];
                            if(valid(newi, newj) && !vis2[newi][newj] && arr[newi][newj] == 0){
                                vis2[newi][newj] = true;
                                q.add(new int[]{newi, newj});
                            }
                        }
                    }

                    if(q.size() > size){
                        size = q.size();
                        ci = i; 
                        cj = j;
                    }

                    q.clear();
                }
            }
        }

        return new int[]{ci, cj};
    }
}