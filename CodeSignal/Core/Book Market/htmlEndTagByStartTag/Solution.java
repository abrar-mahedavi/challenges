public class Solution {

  String solution(String startTag) {
    String[] split = startTag.split("[< >]");
    return "</" + split[1] + ">";
  }

}
