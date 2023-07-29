public class Solution {

  String[] solution(char number) {
    String[] list = new String[(number - 'A') / 2 + 1];
    for (int i = 0 ; i < list.length; i++) {
      list[i] = ((char)(i + 'A')) + " + " + ((char)(number - i));
    }
    return list;
  }

}
