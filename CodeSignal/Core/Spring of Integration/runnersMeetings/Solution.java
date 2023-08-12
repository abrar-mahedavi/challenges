public class Solution {

  int solution(int[] startPosition, int[] speed) {
    int best = -1;
    for (int i = 1; i < speed.length; i++) {
      for (int j = 0; j < i; j++) {
        int m = speed[i] - speed[j];
        if (m != 0) {
          int meet = 2;
          int b = startPosition[j] - startPosition[i];
          for (int k = i+1; k < speed.length; k++) {
            if (startPosition[k]*m + speed[k]*b ==
              startPosition[i]*m + speed[i]*b)
              meet++;
          }
          if (meet > best)
            best = meet;
        }
      }
    }
    return best;
  }
}
