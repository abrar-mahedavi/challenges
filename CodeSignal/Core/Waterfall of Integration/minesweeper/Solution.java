public class Solution {

  int[][] solution(boolean[][] matrix) {
    int[][] mines = new int[matrix.length][matrix[0].length];
    for (int i = 0; i < matrix.length; i++) {
      for (int j = 0; j < matrix[0].length; j++) {
        int sum = 0;
        for (int di = -1; di <= 1; di++) {
          for (int dj = -1; dj <= 1; dj++) {
            if ((di != 0 || dj != 0) &&
              0 <= i + di && i + di < matrix.length &&
              0 <= j + dj && j + dj < matrix[0].length) {
              sum += (matrix[i + di][j + dj] ? 1 : 0);
            }
          }
        }
        mines[i][j] = sum;
      }
    }
    return mines;
  }
}
