public class Solution {

  int solution(int n, int m) {
    int max = 1;
    for (int k = 1; k <= n; k++) {
      if (n%k == 0 && m%k == 0)
        max = k;
    }
    return n+m+max-2;
  }

}
