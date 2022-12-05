public class Solution {

  public static void countTrailingZeros(Integer n) {
    Integer answer = 0;
    for (int i = 5; i <= n; i = i * 5) {
      answer += n / i;
    }
    System.out.println(answer);
  }

  public static void main(String[] args) {
    countTrailingZeros(1);
    countTrailingZeros(5);
    countTrailingZeros(24);
    countTrailingZeros(25);
    countTrailingZeros(26);
    countTrailingZeros(50);
    countTrailingZeros(500);
  }


}
