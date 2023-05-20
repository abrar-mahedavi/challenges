public class Solution {

  int solution(int a) {
    int b = 0;
    while (a > 0) {
      b *= 2;
      b += a%2;
      a /= 2;
    }
    return b;
  }

}
