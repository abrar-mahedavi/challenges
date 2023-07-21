public class Solution {

  int solution(int[] votes, int k) {
    int max = Arrays.stream(votes).max().getAsInt();
    if (k == 0) return Arrays.stream(votes).filter(v -> v == max).count() == 1 ? 1 : 0;
    return (int) Arrays.stream(votes).filter(v -> v + k > max).count();
  }


}
