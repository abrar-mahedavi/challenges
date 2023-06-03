public class Solution {

  int solution(int candlesNumber, int makeNew) {
    return candlesNumber + (candlesNumber-1)/(makeNew-1);
  }

}
