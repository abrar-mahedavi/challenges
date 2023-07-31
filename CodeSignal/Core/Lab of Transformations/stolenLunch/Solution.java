public class Solution {

  String solution(String note) {
    StringBuilder sb = new StringBuilder();
    for (char c : note.toCharArray()) {
      if (c >= '0' && c <= '9') {
        sb.append((char) (c - '0' + 'a'));
      } else if (c >= 'a' && c <= 'j') {
        sb.append((char) (c - 'a' + '0'));
      } else {
        sb.append(c);
      }
    }
    return sb.toString();
  }

}
