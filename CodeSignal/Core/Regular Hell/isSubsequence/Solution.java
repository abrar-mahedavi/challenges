public class Solution {

  boolean solution(String t, String s) {
    String pattern = "";
    for (int i = 0; i < s.length(); i++) {
      pattern += "["+s.charAt(i)+"].*" ;
    }
    Pattern regex = Pattern.compile(pattern);
    return regex.matcher(t).find();
  }

}
