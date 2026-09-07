class Solution {

    // Helper method to simulate string typing and backspace handling 🗂️
    public String eval(String s) {
        Stack<Character> stk = new Stack<>();
        
        for (char ch : s.toCharArray()) {
            if (ch != '#') {
                stk.push(ch); // Type a normal character
            } else if (!stk.isEmpty()) {
                stk.pop();    // Process backspace by removing the last typed character ✂️
            }
        }
        
        // Reconstruct the final string from the remaining stack elements
        StringBuilder res = new StringBuilder();
        for (char ch : stk) {
            res.append(ch);
        }
        return res.toString();
    }

    public boolean backspaceCompare(String s, String t) {
        // Evaluate both strings and verify if their final configurations are identical
        return eval(s).equals(eval(t));
    }
}