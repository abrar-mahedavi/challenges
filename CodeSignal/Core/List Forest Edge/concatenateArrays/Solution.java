public class Solution {

  int[] solution(int[] a, int[] b) {
    return IntStream.concat(Arrays.stream(a), Arrays.stream(b)).toArray();
  }

}
