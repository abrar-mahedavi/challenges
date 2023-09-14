public class Solution {

  String[] solution(String bishop1, String bishop2) {
    String[] result = new String[2];

    int xDiff = bishop1.charAt(0) - bishop2.charAt(0);
    int yDiff = bishop1.charAt(1) - bishop2.charAt(1);
    boolean isSameDiagonal = Math.abs(xDiff) == Math.abs(yDiff);
    if (isSameDiagonal) {
      int x = Integer.signum(xDiff);
      int y = Integer.signum(yDiff);
      result[0] = getEdgeCell(bishop1, x, y);
      result[1] = getEdgeCell(bishop1, -x, -y);
    } else {
      result = new String[]{bishop1, bishop2};
    }

    if (result[0].compareTo(result[1]) > 0) {
      swap(result);
    }
    return result;
  }

  private static void swap(String[] result) {
    String t = result[0];
    result[0] = result[1];
    result[1] = t;
  }

  private static String getEdgeCell(String bishop, int x, int y) {
    int x0 = bishop.charAt(0) - 'a';
    int y0 = bishop.charAt(1) - '1';
    while (x0 + x >= 0 && y0 + y >= 0 && x0 + x <= 7 && y0 + y <= 7) {
      x0 += x;
      y0 += y;
    }
    return (char) (x0 + 'a') + "" + (char) (y0 + '1');
  }


}
