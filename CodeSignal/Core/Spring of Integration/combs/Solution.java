public class Solution {

  int solution(String comb1, String comb2) {
    int bestFit = comb1.length() + comb2.length();
    for (int i = -comb2.length(); i <= comb1.length(); ++i) {
      boolean doesFit = true;
      for (int j = 0; j < comb2.length(); ++j) {
        int k = i + j;
        if (k >= 0 && k < comb1.length() &&
          comb1.charAt(k) == '*' &&
          comb2.charAt(j) == '*') {
          doesFit = false;
        }
      }
      if (!doesFit) { continue; }
      int width = Math.max(comb1.length(), i + comb2.length()) - Math.min(i, 0);
      if (width < bestFit) { bestFit = width; }
    }
    return bestFit;
  }

}
