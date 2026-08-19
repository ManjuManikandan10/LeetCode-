class Solution {
    public int[][] floodFill(int[][] image, int sr, int sc, int color) {
        int originalColor = image[sr][sc];

        // Base case: if target color is same as starting color, return directly
        if (originalColor == color) return image;

        dfs(image, sr, sc, originalColor, color);

        return image;
    }

    void dfs(int[][] image, int sr, int sc, int originalColor, int givenColor) {
        int m = image.length;
        int n = image[0].length;

        // Boundary checks and color match check
        if (sr < 0 || sr >= m || sc < 0 || sc >= n) return;
        if (image[sr][sc] != originalColor) return;

        // Update color
        image[sr][sc] = givenColor;

        // Recurse on adjacent pixels using originalColor
        dfs(image, sr + 1, sc, originalColor, givenColor);
        dfs(image, sr - 1, sc, originalColor, givenColor);
        dfs(image, sr, sc + 1, originalColor, givenColor);
        dfs(image, sr, sc - 1, originalColor, givenColor);
    }
}