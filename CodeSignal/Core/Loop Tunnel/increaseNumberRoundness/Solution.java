public class Solution {

  boolean solution(int n) {
    while (n%10 == 0 && n > 0)
      n /= 10;
    while (n%10 != 0 && n > 0)
      n /= 10;
    return (n != 0);
  }

}
