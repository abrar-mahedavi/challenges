public class Solution {

  boolean solution(String inputString) {
    return inputString.matches("([0-9A-F]{2}-){5}[0-9A-F]{2}");
  }

}
