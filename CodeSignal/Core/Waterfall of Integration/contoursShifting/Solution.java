public class Solution {

  int[][] solution(int[][] matrix) {
    for (int i = 0; matrix.length - i > i && matrix[0].length - i > i; ++i) {
      int br = matrix.length - 1 - i, rc = matrix[0].length - 1 - i; //bottom row, right column
      if ((i & 1) == 1) {
        int aux = matrix[i][i];
        for (int c = i; c < rc; ++c) {
          matrix[i][c] = matrix[i][c + 1];
        }
        for (int r = i; r < br; ++r) {
          matrix[r][rc] = matrix[r + 1][rc];
        }
        if (i < br) {
          for (int c = rc; c > i; --c) {
            matrix[br][c] = matrix[br][c - 1];
          }
        }
        if (i < rc) {
          for (int r = br; r > i; --r) {
            matrix[r][i] = matrix[r - 1][i];
          }
        }
        if (i < rc && i < br) {
          matrix[i + 1][i] = aux;
        } else {
          matrix[br][rc] = aux;
        }
      } else {
        int aux = matrix[br][rc];
        for (int r = br; r > i; --r) {
          matrix[r][rc] = matrix[r - 1][rc];
        }
        for (int c = rc; c > i; --c) {
          matrix[i][c] = matrix[i][c - 1];
        }
        if (i < rc) {
          for (int r = i; r < br; ++r) {
            matrix[r][i] = matrix[r + 1][i];
          }
        }
        if (i < br) {
          for (int c = i; c < rc; ++c) {
            matrix[br][c] = matrix[br][c + 1];
          }
        }
        if (i < rc && i < br) {
          matrix[br][rc - 1] = aux;
        } else {
          matrix[i][i] = aux;
        }
      }
    }
    return matrix;
  }


}
