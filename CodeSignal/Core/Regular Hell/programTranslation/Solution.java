public class Solution {

  String solution(String code, String[] args) {
    String argumentVariants = String.join("|", args);
    String pattern = "(?<![\\$\\w])("+argumentVariants+")(?![\\d\\w])" ;
    String sub = "\\$$1" ;
    return code.replaceAll(pattern, sub);
  }

}
