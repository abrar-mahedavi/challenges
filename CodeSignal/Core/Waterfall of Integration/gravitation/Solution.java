public class Solution {

  int[] solution(String[] rows) {
    int[] t = new int[rows[0].length()];
    for (int i = 0; i < rows[0].length(); i++) {

      int j = 0;
      for (; j < rows.length; j++) {
        if (rows[j].charAt(i) == '#') {
          break;
        }
      }

      for (; j < rows.length; j++) {
        if (rows[j].charAt(i) == '.') {
          t[i]++;
        }
      }
    }

    int min = 1000;
    int c = 0;
    for (int i = 0; i < t.length; i++) {
      if (t[i] < min) {
        min = t[i];
        c = 1;
      } else if (t[i] == min) {
        c++;
      }
    }
    int[] res = new int[c];
    int i = 0;
    for (int k = 0; k < t.length; k++) {
      if (t[k] == min) {
        res[i++] = k;
      }
    }

    return res;
  }


}
