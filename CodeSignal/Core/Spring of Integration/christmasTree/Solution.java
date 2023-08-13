public class Solution {

  String[] solution(int levelNum, int levelHeight) {
    final char newline = '\n';
    final int maxWidth = 2 * (levelNum + levelHeight) + 3;
    final int footWidth = levelHeight % 2 == 0 ? levelHeight + 1 : levelHeight;

    final StringBuilder sb = new StringBuilder();

    for (int i = 1; i < maxWidth / 2; ++i) {
      sb.append(' ');
    }
    sb.append('*').append(newline);

    for (int i = 1; i < maxWidth / 2; ++i) {
      sb.append(' ');
    }
    sb.append('*').append(newline);

    for (int i = 1; i < maxWidth / 2 - 1; ++i) {
      sb.append(' ');
    }
    sb.append("***").append(newline);

    int initialWidth = 5;
    for (int i = 0; i < levelNum; ++i) {            // <= every level
      int rowWidth = initialWidth;
      for (int j = 0; j < levelHeight; ++j) {    // <= every row in level
        for (int k = 1; k < maxWidth; ++k) {    // <= handle row
          if (k < (maxWidth / 2) - (rowWidth / 2)) {
            sb.append(' ');
          } else if (k > (maxWidth / 2) + (rowWidth / 2)) {
            break;
          } else {
            sb.append('*');
          }
        }
        sb.append(newline);
        rowWidth += 2;
      }

      initialWidth += 2;
    }

    for (int i = 0; i < levelNum; ++i) {
      for (int j = 1; j < maxWidth; ++j) {
        if (j < (maxWidth / 2) - (footWidth / 2)) {
          sb.append(' ');
        } else if (j > (maxWidth / 2) + (footWidth / 2)) {
          break;
        } else {
          sb.append('*');
        }
      }
      sb.append(newline);
    }

    return sb.toString().split("\\n");
  }
}
