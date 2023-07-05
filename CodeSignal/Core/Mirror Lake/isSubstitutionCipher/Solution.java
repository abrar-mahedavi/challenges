public class Solution {

  boolean solution(String string1, String string2) {
    if(string1.length() != string2.length())return false;
    String t1= "", t2= "";
    for(int i= 0;i < string1.length();++i){
      int iof1= t1.indexOf(string1.charAt(i));
      int iof2= t2.indexOf(string2.charAt(i));
      if(iof1 != iof2)return false;
      if(iof1 < 0){
        t1+= string1.charAt(i);
        t2+= string2.charAt(i);
      }
    }
    return true;
  }

}
