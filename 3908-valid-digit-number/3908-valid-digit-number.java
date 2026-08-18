class Solution {
    public boolean validDigit(int n, int x) {
        boolean occurrenceX = false;

        while (n >= 10) {
            if (n % 10 == x) {
                occurrenceX = true;
            }

            n /= 10;
        }

        return occurrenceX && n != x;
    }
}