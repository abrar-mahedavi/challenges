public class Solution {

  int[] solution(int[][] matrix, int column) {
    return Arrays.stream(matrix).mapToInt(row -> row[column]).toArray();
  }

}
