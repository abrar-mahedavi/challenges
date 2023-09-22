public class Solution {

  int[] solution(String part, String total) {
    int[] a = new int[2] ;
    String[] s1 = part.split(":") ;
    String[] s2 = total.split(":") ;
    int l1 = Integer.valueOf(s1[0])*60*60 + Integer.valueOf(s1[1])*60 + Integer.valueOf(s1[2]) ;
    int l2 = Integer.valueOf(s2[0])*60*60+ Integer.valueOf(s2[1])*60 + Integer.valueOf(s2[2]) ;
    int gcm =(int) gcm(l1, l2);
    a[0] = l1/gcm ;
    a[1] = l2/gcm ;
    return a ;
  }
  int gcm(int  a, int  b) {
    return b == 0 ? a : gcm(b, a % b);
  }
}
