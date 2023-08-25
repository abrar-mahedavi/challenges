public class Solution {

  boolean solution(int[][] grid) {
    ArrayList<Integer> werte = new ArrayList<Integer>();
    for (int i = 0; i < 9; i++) {
      werte.clear();
      for (int j = 0; j < 9; j++) {
        if (!(werte.contains(grid[i][j]))) {
          werte.add(grid[i][j]);
        } else {
          return false;
        }
      }
    }
    for (int i = 0; i < 9; i++) {
      werte.clear();
      for (int j = 0; j < 9; j++) {
        if (!(werte.contains(grid[j][i]))) {
          werte.add(grid[j][i]);
        } else {
          return false;
        }
      }
    }
    for (int y = 0; y < 3; y++) {
      for (int i = 0; i < 3; i++) {
        werte.clear();
        for (int j = 0; j < 3; j++) {

          for (int x = 0; x < 3; x++) {
            if (!(werte.contains(grid[j + 3 * i][x + 3 * y]))) {
              werte.add(grid[j + 3 * i][x + 3 * y]);
              System.out.println(grid[j + 3 * i][x + 3 * y]);
            } else {
              return false;
            }
          }
        }
      }
    }
    return true;
  }


}
