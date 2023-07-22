public class Solution {

  String solution(int number, int width) {
    String s = String.format("%0" + width + "d", number);
    return s.substring(s.length() - width);
  }

}
