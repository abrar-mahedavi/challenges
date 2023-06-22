public class Solution {

  int solution(int a, int b) {
    int x = (int)(a / 2 / Math.sqrt(2));
    int y = (int)(b / 2 / Math.sqrt(2));
    int k = (int)((a / 2 - Math.sqrt(2)  / 2) / Math.sqrt(2));
    int z = (int)((b / 2 - Math.sqrt(2)  / 2) / Math.sqrt(2));
    return 2 * ((x + 1) * (y + 1) + 2 * (k + 1) * (z + 1) + x * y) - 1;
  }

}
