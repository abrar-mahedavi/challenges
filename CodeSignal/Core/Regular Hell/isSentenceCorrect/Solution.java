public class Solution {

  boolean solution(String sentence) {
    String regex = "[A-Z][^.!?]*[.!?]{1}$";
    return sentence.matches(regex);
  }

}
