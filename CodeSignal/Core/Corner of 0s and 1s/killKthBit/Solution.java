public class Solution {

  int solution(int n, int k) {
    return n & ~(1 << (k - 1));
  }

}
