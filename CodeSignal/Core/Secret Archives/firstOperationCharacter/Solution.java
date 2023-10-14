public class Solution {

  int solution(String expr) {
    int bracket = 0;
    int plusMax = -1;
    int multMax = -1;
    int plusIdx = 0;
    int multIdx = 0;
    for (int i = 0; i < expr.length(); i++) {
      char x = expr.charAt(i);
      if (x == '(') bracket++;
      if (x == ')') bracket--;
      if (x == '+') {
        if (plusMax < bracket) {
          plusIdx = i;
          plusMax = bracket;
        }

      }
      if (x == '*') {
        if (multMax < bracket){
          multIdx = i;
          multMax = bracket;
        }
      }
    }

    return (plusMax > multMax) ? plusIdx : multIdx;
  }


}
