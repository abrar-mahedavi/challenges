public class Solution {

  int solution(int[] coins, int price) {
    int t = 0;
    for (int i = coins.length-1; i >= 0; i--) {
      t += price/coins[i];
      price %= coins[i];
    }
    return t;
  }

}
