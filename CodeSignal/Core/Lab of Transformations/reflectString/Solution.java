public class Solution {

  String solution(String inputString) {
    String mirror = "";

    for( int i = 0 ; i < inputString.length() ; i++ ){
      mirror += (char)('z' - (inputString.charAt(i) - 'a'));

    }
    return mirror;
  }

}
