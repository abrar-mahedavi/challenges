public class Solution {

  boolean solution(int[] length, int[] width, int[] height) {
    int[][] a = IntStream.range(0, length.length)
      .mapToObj(i -> Arrays.stream(new int[]{length[i], width[i], height[i]}).sorted().toArray())
      .sorted((x, y) -> x[0] - y[0]).toArray(int[][]::new);
    return IntStream.range(0, 3)
      .mapToObj(j -> IntStream.range(1, a.length).allMatch(i -> a[i - 1][j] < a[i][j]))
      .allMatch(i -> i);
  }


}
