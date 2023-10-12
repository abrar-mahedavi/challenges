public class Solution {

  String solution(String notation) {
    char[][] board = new char[8][8];
    int i = 0;
    int j = 0;
    for(int k = 0; k < notation.length(); k++) {
      char ch = notation.charAt(k);
      if (ch == '/') {
        i++;
        j = 0;
        continue;
      }
      if (Character.isDigit(ch)) {
        j += ch - '0';
        continue;
      }
      board[i][j++] = ch;
    }
    StringBuilder sb = new StringBuilder();
    for(int col = 0; col < board.length; col++) {
      if (col != 0) {
        sb.append('/');
      }
      int n = 0;
      for (int row = board[col].length - 1; row >= 0; row--) {
        char ch = board[row][col];
        if (ch == 0) {
          n++;
        } else {
          if (n != 0) {
            sb.append(n);
            n = 0;
          }
          sb.append(ch);
        }
      }
      if (n != 0) {
        sb.append(n);
      }
    }
    return sb.toString();
  }


}
