public class Solution {

  int[] solution(int[] shuffled) {
    int n = IntStream.of(shuffled).sum() / 2;
    int p = IntStream.range(0, shuffled.length)
      .filter(i -> shuffled[i] == n).findFirst().getAsInt();
    return IntStream.range(0, shuffled.length)
      .filter(i -> i != p)
      .map(i -> shuffled[i])
      .sorted()
      .toArray();
  }

}
