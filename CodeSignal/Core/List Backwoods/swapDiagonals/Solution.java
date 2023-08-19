public class Solution {

  int[][] solution(int[][] matrix) {
    for (int i = 0; i < matrix.length; i++) {
      int j = matrix.length - i - 1;
      int temp = matrix[i][i];
      matrix[i][i] = matrix[i][j];
      matrix[i][j] = temp;
    }
    return matrix;
  }

}
