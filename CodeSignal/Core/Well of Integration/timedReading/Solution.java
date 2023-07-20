public class Solution {

  int solution(int maxLength, String text) {
    return (int) Arrays.stream(text.split("[^A-z]+")).filter(s -> s.length() <= maxLength).count();
  }

}
