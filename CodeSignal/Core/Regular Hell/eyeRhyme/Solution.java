public class Solution {

  Boolean solution(String pairOfLines) {
    Pattern pattern = Pattern.compile(".*(...)\t.*(...)");
    Matcher matcher = pattern.matcher(pairOfLines);
    matcher.matches();
    return matcher.group(1).equals(matcher.group(2));
  }

}
