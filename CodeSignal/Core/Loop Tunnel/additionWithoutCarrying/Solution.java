public class Solution {

  int solution(int param1, int param2) {
    int t = 0;
    int mult = 1;
    while (param1 + param2 > 0) {
      int value = (param1 +param2)%10;
      t += value * mult;
      mult *= 10;
      param1 /= 10;
      param2 /= 10;
    }
    return t;
  }
}
