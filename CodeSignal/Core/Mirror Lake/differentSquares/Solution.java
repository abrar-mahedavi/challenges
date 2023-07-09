public class Solution {

  int solution(int[][] matrix) {
    HashSet s = new HashSet();
    for (int i = 0; i < matrix.length - 1; i++) {
      for (int j = 0; j < matrix[i].length - 1; j++) {
        s.add(matrix[i][j] + "," + matrix[i][j+1] + "," + matrix[i+1][j] + "," + matrix[i+1][j+1]);}}
    return s.size();
  }

}
