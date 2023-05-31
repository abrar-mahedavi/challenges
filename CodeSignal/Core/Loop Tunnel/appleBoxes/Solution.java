public class Solution {

  int solution(int k) {
    int t = 0;
    for (int i = 1; i <= k; i++) {
      if (i%2 == 1)
        t -= i*i;
      else
        t += i*i;
    }
    return t;
  }

}
