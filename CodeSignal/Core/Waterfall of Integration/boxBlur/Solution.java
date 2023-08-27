public class Solution {

  int[][] solution(int[][] image) {
    int[][] rs = new int[image.length - 2][image[0].length - 2];
    for (int i = 1; i < image.length - 1; i++) {
      for (int j = 1; j < image[i].length - 1; j++) {
        int avg = image[i][j] + image[i - 1][j - 1] + image[i - 1][j]
          + image[i - 1][j + 1] + image[i][j + 1] + image[i + 1][j + 1]
          + image[i + 1][j] + image[i + 1][j - 1] + image[i][j - 1];
        rs[i - 1][j - 1] = avg / 9;
      }
    }
    return rs;

  }

}
