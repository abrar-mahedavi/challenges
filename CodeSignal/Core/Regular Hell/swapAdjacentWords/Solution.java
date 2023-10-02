public class Solution {

  String solution(String s) {
    return s.replaceAll("(\\w+) (\\w+)", "$2 $1");
  }

}
