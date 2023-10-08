public class Solution {

  int solution(String rules) {
    Pattern pattern = Pattern.compile("(\\d*)d(\\d+)([\\+\\-]\\d+)?");
    Matcher m = pattern.matcher(rules);

    int res = 0;
    while (m.find()) {
      int rolls = m.group(1).isEmpty() ? 1 : Integer.parseInt(m.group(1));
      int dieType = Integer.parseInt(m.group(2));
      int formulaMax = rolls * dieType;

      if (m.group(3) != null && !m.group(3).isEmpty()) {
        if (m.group(3).charAt(0) == '-') {
          formulaMax -= Integer.parseInt(m.group(3).substring(1));
        }
        else {
          formulaMax += Integer.parseInt(m.group(3).substring(1));
        }
      }

      res += formulaMax;
    }

    return res;
  }

}
