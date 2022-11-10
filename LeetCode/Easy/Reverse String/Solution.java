// Online Java Compiler
// Use this editor to write, compile and run your Java code online

class Solution {
    public static void main(String[] args) {
        System.out.println("Hello, World!");
        Solution solution = new Solution();
        solution.reverseString(new char[] { '1', '2', '1', '2', '3' });
    }

    public void reverseString(char[] s) {
        for (int i = 0; i < s.length / 2; i++) {
            char temp = s[i];
            s[i] = s[s.length - i - 1];
            s[s.length - i - 1] = temp;
        }

    }

}