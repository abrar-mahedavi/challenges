public class Solution {

  int solution(char[][][] pieces) {
    int[][] map = new int[20][10];
    int total = 0;
    int counter = 1;
    for (char[][] piece : pieces) {
      int best = -1;
      int bestX = -1;
      int bestD = -1;
      for (int dir = 0; dir < 4; dir++) {
        for (int i = 0; i < 10; i++) {
          int value = score(map, piece, i);
          if (value > best) {
            best = value;
            bestX = i;
            bestD = dir;
          }
        }
        piece = rotate(piece);
      }
      System.out.println(bestX + " : " + bestD + " : " + counter);
      map = place(map, piece, bestX, bestD, counter);
      int rows = check(map);
      if (rows > 0) {
        map = fix(map);
      }
      total += rows;
      counter++;
      for (int i = 0; i < 20; i++) {
        for (int j = 0; j < 10; j++) {
          System.out.print(map[i][j] + " ");
        }
        System.out.println();
      }
      System.out.println();
    }
    return total;
  }

  char[][] rotate(char[][] piece) {
    char[][] result = new char[piece[0].length][piece.length];
    for (int i = 0; i < piece.length; i++) {
      for (int j = 0; j < piece[0].length; j++) {
        result[j][piece.length - 1 - i] = piece[i][j];
      }
    }
    return result;
  }

  int score(int[][] map, char[][] piece, int x) {
    int loc = 20 - piece.length;
    if (x + piece[0].length > 10) {
      return -1;
    }
    for (int i = 0; i < 21 - piece.length; i++) {
      for (int j = 0; j < piece.length; j++) {
        for (int k = 0; k < piece[0].length; k++) {
          if (map[i + j][x + k] > 0 && piece[j][k] == '#') {
            j = 5;
            k = 5;
            loc = i - 1;
            i = 21;
          }
        }
      }
    }
    if (loc < 0) {
      return -1;
    }
    int total = 0;
    for (int i = loc; i < loc + piece.length; i++) {
      for (int j : map[i]) {
        if (j > 0) {
          total++;
        }
      }
    }
    return total;
  }

  int[][] place(int[][] map, char[][] piece, int x, int dir, int val) {
    if (val < 1) {
      val = 1;
    }
    for (int i = 0; i < dir; i++) {
      piece = rotate(piece);
    }
    int loc = 20 - piece.length;
    for (int i = 0; i < 21 - piece.length; i++) {
      for (int j = 0; j < piece.length; j++) {
        for (int k = 0; k < piece[0].length; k++) {
          if (map[i + j][x + k] > 0 && piece[j][k] == '#') {
            k = 5;
            j = 5;
            loc = i - 1;
            i = 21;
          }
        }
      }
    }
    for (int j = 0; j < piece.length; j++) {
      for (int k = 0; k < piece[0].length; k++) {
        if (piece[j][k] == '#') {
          map[loc + j][x + k] = val;
        }
      }
    }
    return map;
  }

  int check(int[][] map) {
    int total = 0;
    for (int i = 0; i < 20; i++) {
      boolean full = true;
      for (int j : map[i]) {
        if (j == 0) {
          full = false;
        }
      }
      if (full) {
        total++;
      }
    }
    return total;
  }

  int[][] fix(int[][] map) {
    for (int i = 0; i < 20; i++) {
      boolean full = true;
      for (int j : map[i]) {
        if (j == 0) {
          full = false;
        }
      }
      if (full) {
        for (int j = i; j > 0; j--) {
          for (int k = 0; k < 10; k++) {
            map[j][k] = map[j - 1][k];
            map[j - 1][k] = 0;
          }
        }
      }
    }
    return map;
  }

}
