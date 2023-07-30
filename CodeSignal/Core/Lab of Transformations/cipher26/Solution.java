public class Solution {

  String solution(String message) {
    String a = "abcdefghijklmnopqrstuvwxyz", s = "";
    int n = 0, t = 0;
    for (int i = 0; i < message.length(); i++) {
      t = (((a.indexOf(message.substring(i, i + 1)) - n) % 26) + 26) % 26;
      s += a.substring(t, t + 1);
      n += t;
    }
    return s;
  }
}
