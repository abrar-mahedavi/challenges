public class Solution {

  int solution(int n, int m) {
    return
      8 * m(n - 1) * m(m - 2)
        + 8 * m(n - 1) * m(m - 3)
        + 8 * m(n - 2) * m(m - 1)
        + 16 * m(n - 2) * m(m - 2)
        + 8 * m(n - 2) * m(m - 3)
        + 8 * m(n - 3) * m(m - 1)
        + 8 * m(n - 3) * m(m - 2);
  }

  static int m(int n) {
    return Math.max(n, 0);
  }


}
