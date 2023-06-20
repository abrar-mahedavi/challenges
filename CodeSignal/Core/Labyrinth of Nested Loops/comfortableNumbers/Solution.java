public class Solution {

  int solution(int l, int r) {
    int total = 0;
    for (int i = l; i <= r; i++) {
      for (int j = i + 1; j <= r; j++) {
        int sumi = String.valueOf(i).chars().map(a -> Character.getNumericValue(a)).sum();
        int sumj = String.valueOf(j).chars().map(a -> Character.getNumericValue(a)).sum();
        if (j >= (i - sumi) && j <= (i + sumi) && i >= (j - sumj) && i <= (j + sumj)) {
          total++;
        }
      }
    }
    return total;
  }

}
