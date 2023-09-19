public class Solution {

  String solution(String white, String black, char toMove) {
    int[] w = chessToInt(white);
    int[] b = chessToInt(black);
    if (w[0] == b[0] && w[1] < b[1]) {
      return "draw";
    }
    return move(w[0], w[1], b[0], b[1], toMove == 'w') ? "white" : "black";
  }

  int[] chessToInt(String chess) {
    int[] ret = new int[2];
    ret[0] = chess.charAt(0) - 'a';
    ret[1] = chess.charAt(1) - '1';
    return ret;
  }

  boolean move(int w1, int w2, int b1, int b2, boolean wmove) {
    System.out.println("" + w1 + w2 + b1 + b2);
    if (w2 == b2 - 1 && (w1 == b1 + 1 || w1 == b1 - 1)) {
      return wmove;
    }
    if (wmove) {
      if (b2 == 0) {
        return false;
      }
      if (w2 == 1) {
        return move(w1, w2 + 1, b1, b2, false) || move(w1, w2 + 2, b1, b2, false);
      } else {
        return move(w1, w2 + 1, b1, b2, false);
      }
    } else {
      if (w2 == 7) {
        return true;
      }
      if (b2 == 6) {
        return move(w1, w2, b1, b2 - 1, true) && move(w1, w2, b1, b2 - 2, true);
      } else {
        return move(w1, w2, b1, b2 - 1, true);
      }
    }
  }

}
