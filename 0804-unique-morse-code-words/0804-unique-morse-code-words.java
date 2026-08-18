class Solution {
    public int uniqueMorseRepresentations(String[] words) {
        String[] morse = {".-","-...","-.-.","-..",".","..-.","--.","....","..",".---","-.-",".-..","--","-.","---",".--.","--.-",".-.","...","-","..-","...-",".--","-..-","-.--","--.."};
        
        Set<String> transformations = new HashSet<>();
        for (String word : words) {
            StringBuilder current = new StringBuilder();
            for (char ch : word.toCharArray()) {
                current.append(morse[ch - 'a']);
            }
            transformations.add(current.toString());
        }
        return transformations.size();
    }
}