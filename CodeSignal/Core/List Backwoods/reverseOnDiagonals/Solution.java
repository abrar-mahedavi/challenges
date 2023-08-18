public class Solution {

  int[][] solution(int[][] matrix) {
    int l = matrix.length;
    for (int i = 0; i < l / 2; i++) {
      int j = l - 1 - i;
      int t = matrix[i][i];
      matrix[i][i] = matrix[j][j];
      matrix[j][j] = t;
      t = matrix[j][i];
      matrix[j][i] = matrix[i][j];
      matrix[i][j] = t;
    }
    return matrix;
  }

}
