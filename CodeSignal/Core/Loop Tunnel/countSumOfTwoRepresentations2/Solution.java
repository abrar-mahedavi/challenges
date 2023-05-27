public class Solution {

  int solution(int n, int l, int r) {
    int t = 0;
    for (int i = l; i <= r; i++) {
      if (n-i >= i && n-i <= r)
        t++;
    }
    return t;
  }
}
