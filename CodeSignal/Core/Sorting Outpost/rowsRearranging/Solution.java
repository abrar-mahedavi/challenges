public class Solution {

  boolean solution(int[][] matrix) {
    Arrays.sort(matrix, (x, y) -> x[0] - y[0]);
    return IntStream.range(0, matrix[0].length).mapToObj(
        j -> IntStream.range(1, matrix.length).allMatch(i -> matrix[i - 1][j] < matrix[i][j]))
      .allMatch(i -> i);
  }

}
