public class Solution {

  int solution(String birthdayDate) {
    String A[] = birthdayDate.split("-");
    int d = Integer.parseInt(A[1]), m = Integer.parseInt(A[0]), y = Integer.parseInt(A[2]);
    if (m == 2 && d == 29) {
      int d2 = zeller(y, m, d);
      for (int i = y + 4; ; i += 4) {
        if (isleap(i)) {
          int d3 = zeller(i, m, d);
          if (d2 == d3) {
            return i - y;
          }
        }
      }
    } else {
      int d2 = zeller(y, m, d), d3;
      for (int i = y + 1; ; i++) {
        d3 = zeller(i, m, d);
        if (d2 == d3) {
          return i - y;
        }
      }
    }

  }

  boolean isleap(int y) {
    return (y % 4 == 0 && (y % 100 != 0 || y % 400 == 0));
  }

  int zeller(int y, int m, int d) {
    if (m < 3) {
      --y;
    }
    m += 12;
    return (y + y / 4 - y / 100 + y / 400 + (13 * m + 8) / 5 + d) % 7;
  }


}
