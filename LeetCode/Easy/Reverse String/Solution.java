// Online Java Compiler
// Use this editor to write, compile and run your Java code online

class Solution {
    public static void main(String[] args) {
        System.out.println("Hello, World!");
        Solution solution = new Solution();
        int output = solution.singleNumber(new int[] { 1, 2, 1, 2, 3 });
        System.out.println(output);
    }

    public int singleNumber(int[] nums) {
        if (nums.length == 1) {
            return nums[0];
        } else {
            int sum = 0;
            for (int i = 0; i < nums.length; i++) {
                sum ^= nums[i];
            }
            return sum;
        }
    }
}