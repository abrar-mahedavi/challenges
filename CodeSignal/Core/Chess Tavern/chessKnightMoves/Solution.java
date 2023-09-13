public class Solution {

  int solution(String cell) {
    int x = cell.charAt(0) % 97,
      y = +cell.charAt(1) - '0' - 1;
    int c = 0;
    for (int dx = -2; dx <= 2; dx++) {
      for (int dy = -2; dy <= 2; dy++) {
        if (Math.abs(dx * dy) == 2) {
          if (isValid(x + dx, y + dy)) {
            c++;
          }
        }
      }
    }
    return c;
  }

  private static boolean isValid(int x, int y) {
    return x >= 0 && x <= 7 && y >= 0 && y <= 7;
  }

}
