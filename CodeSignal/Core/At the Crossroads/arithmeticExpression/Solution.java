public class Solution {

  boolean solution(int a, int b, int c) {
    return a + b == c || a - b == c || a * b == c || a / (b * 1.0) == c;
  }

}
