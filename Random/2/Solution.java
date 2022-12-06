public class Solution {

  public static void gcd(Integer a,Integer b){
    while (a!=0 && b!=0){
      if (a > b ){
        a = a%b;
      }else {
        b = b%a;
      }
    }
    System.out.println(a > 0 ? a: b );
  }

  public static void main(String[] args) {
    gcd(1,2);
    gcd(2,4);
    gcd(24,36);
    gcd(25,50);
    gcd(38,54);
    gcd(171,70);
    gcd(500,500);
  }


}
