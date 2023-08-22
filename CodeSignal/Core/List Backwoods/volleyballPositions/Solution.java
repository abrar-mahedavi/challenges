public class Solution {

  String[][] solution(String[][] formation, int k) {
    for (int i = k % 6; i > 0; --i) {
      final String temp = formation[0][1];

      formation[0][1] = formation[1][2];
      formation[1][2] = formation[3][2];
      formation[3][2] = formation[2][1];
      formation[2][1] = formation[3][0];
      formation[3][0] = formation[1][0];
      formation[1][0] = temp;
    }

    return formation;
  }

}
