import java.util.Stack;

// Online Java Compiler
// Use this editor to write, compile and run your Java code online

class Solution {
    public static void main(String[] args) {
        System.out.println("Hello, World!");
        Solution solution = new Solution();
        System.out.println(solution.isValid(new char[] { '(', '{', '}', ')' }));
        System.out.println(solution.isValid("({})"));
    }

    public static boolean isValid(String s) {
        if (s.length() % 2 != 0)
            return false;
        Stack<Character> stack = new Stack<Character>();
        for (char c : s.toCharArray()) {
            switch (c) {
                case '(':
                    stack.push(')');
                    break;
                case '{':
                    stack.push('}');
                    break;
                case '[':
                    stack.push(']');
                    break;
                case ')':
                case '}':
                case ']':
                    if (stack.isEmpty() || stack.pop() != c)
                        return false;
            }
        }
        return stack.isEmpty();
    }

    public static boolean isValid(char[] s) {
        if (s.length % 2 != 0)
            return false;
        Stack<Character> stack = new Stack<Character>();
        for (int i = 0; i < s.length; i++) {
            char c = s[i];
            switch (c) {
                case '(':
                    stack.push(')');
                    break;
                case '{':
                    stack.push('}');
                    break;
                case '[':
                    stack.push(']');
                    break;
                case ')':
                case '}':
                case ']':
                    if (stack.isEmpty() || stack.pop() != c)
                        return false;
            }
        }
        return stack.isEmpty();
    }
}