public class Solution {

  int solution(String inputString) {
    int count = 0;
    Matcher matcher = Pattern.compile("\\d+|\"[^\"]*\"|true|false").matcher(inputString);
    while (matcher.find()) {
      count++;
    }
    return count;
  }

}
