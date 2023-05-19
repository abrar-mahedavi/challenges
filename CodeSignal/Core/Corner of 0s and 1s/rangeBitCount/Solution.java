public class Solution {

  int solution(int a, int b) {
    int t = 0;
    for (int i = a; i<=b; i++) {
      t += Integer.bitCount(i);
    }
    return t;
  }
}
