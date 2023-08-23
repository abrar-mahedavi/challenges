public class Solution {

  int[][] solution(int[][] matrix, int width, int[] center, int t) {
    final int centerX = center[1];
    final int centerY = center[0];

    for (int i = t % 8; i > 0; --i) {
      for (int j = width / 2; j > 0; --j) {
        final int temp = matrix[centerY - j][centerX - j];
        matrix[centerY - j][centerX - j] = matrix[centerY][centerX - j];
        matrix[centerY][centerX - j] = matrix[centerY + j][centerX - j];
        matrix[centerY + j][centerX - j] = matrix[centerY + j][centerX];
        matrix[centerY + j][centerX] = matrix[centerY + j][centerX + j];
        matrix[centerY + j][centerX + j] = matrix[centerY][centerX + j];
        matrix[centerY][centerX + j] = matrix[centerY - j][centerX + j];
        matrix[centerY - j][centerX + j] = matrix[centerY - j][centerX];
        matrix[centerY - j][centerX] = temp;
      }
    }

    return matrix;
  }

}
