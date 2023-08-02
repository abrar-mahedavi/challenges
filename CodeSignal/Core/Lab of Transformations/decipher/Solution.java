public class Solution {

  String solution(String cipher) {
    String ret = "";

    String temp = "";

    for(int i = 0; i<cipher.length(); i++){
      temp += cipher.charAt(i);
      int j = Integer.parseInt(temp);

      if(j>=97 && j<=122){
        ret += (char)j;
        temp = "";
      }
    }

    return ret;
  }

}
