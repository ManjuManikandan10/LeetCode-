class Solution {
    public boolean isLongPressedName(String name, String typed) {
        // Initialize pointers for name (i) and typed (j)
        int i = 0, n = name.length(), m = typed.length();

        // Loop through each character in 'typed'
        for (int j = 0; j < m; ++j) {
            // If current chars match, move both pointers
            if (i < n && name.charAt(i) == typed.charAt(j)) {
                ++i;
            } 
            // If no match, check if it's a long-pressed repeat
            else if (j == 0 || typed.charAt(j) != typed.charAt(j - 1)) {
                return false; // Not a valid long-pressed name
            }
        }

        // Check if we've traversed the entire 'name'
        return i == n;
    }
}