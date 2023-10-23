public class Solution {

  char[][] solution(char[] faceColors, String[] moves) {

    String[] rMoves = moves.clone();
    for (int i = 0; i < rMoves.length; i++) {
      if (rMoves[i].charAt(rMoves[i].length() - 1) == '\'') {
        rMoves[i] = rMoves[i].substring(0, rMoves[i].length() - 1);
      } else {
        rMoves[i] += "'";
      }
    }
    Pyraminx pyra = new Pyraminx(faceColors);
    for (int i = rMoves.length - 1, j = 0; i >= 0; i--, j++) {
      pyra.input(rMoves[i]);
    }
    return pyra.faces;
  }


  public class Pyraminx {

    Pyraminx(char[] faceColors) {
      faces = new char[4][9];
      for (int i = 0; i < 4; i++) {
        Arrays.fill(faces[i], faceColors[i]);
      }
    }

    public char[][] faces;
    public int FRONT = 0, BACK = 1, LEFT = 2, RIGHT = 3;

    public void input(String command) {
      int[] cwTipMap = {0, 8, 4};
      int[] ccwTipMap = {0, 4, 8};
      int[][] ccwRowMap = {{1, 2, 3}, {6, 5, 1}, {3, 7, 6}};
      int[][] cwRowMap = {{1, 2, 3}, {3, 7, 6}, {6, 5, 1}};
      switch (command) {
        case "u":
          rotRow(FRONT, LEFT, RIGHT, ccwRowMap);
        case "U":
          rotTip(FRONT, LEFT, RIGHT, ccwTipMap);
          break;
        case "u'":
          rotRow(FRONT, RIGHT, LEFT, cwRowMap);
        case "U'":
          rotTip(FRONT, RIGHT, LEFT, cwTipMap);
          break;
        case "b":
          rotRow(BACK, RIGHT, LEFT, ccwRowMap);
        case "B":
          rotTip(BACK, RIGHT, LEFT, ccwTipMap);
          break;
        case "b'":
          rotRow(BACK, LEFT, RIGHT, cwRowMap);
        case "B'":
          rotTip(BACK, LEFT, RIGHT, cwTipMap);
          break;
        case "l":
          rotRow(LEFT, FRONT, BACK, ccwRowMap);
        case "L":
          rotTip(LEFT, FRONT, BACK, ccwTipMap);
          break;
        case "l'":
          rotRow(LEFT, BACK, FRONT, cwRowMap);
        case "L'":
          rotTip(LEFT, BACK, FRONT, cwTipMap);
          break;
        case "r":
          rotRow(RIGHT, BACK, FRONT, ccwRowMap);
        case "R":
          rotTip(RIGHT, BACK, FRONT, ccwTipMap);
          break;
        case "r'":
          rotRow(RIGHT, FRONT, BACK, cwRowMap);
        case "R'":
          rotTip(RIGHT, FRONT, BACK, cwTipMap);
          break;
      }
    }

    private void rotTip(int a, int b, int c, int[] d) {
      char tmp = faces[a][d[0]];
      faces[a][d[0]] = faces[b][d[1]];
      faces[b][d[1]] = faces[c][d[2]];
      faces[c][d[2]] = tmp;

    }

    private void rotRow(int a, int b, int c, int[][] d) {
      for (int i = 0; i < 3; i++) {
        char tmp = faces[a][d[0][i]];
        faces[a][d[0][i]] = faces[b][d[1][i]];
        faces[b][d[1][i]] = faces[c][d[2][i]];
        faces[c][d[2][i]] = tmp;
      }
    }
  }

}
