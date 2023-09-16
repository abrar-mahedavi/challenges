public class Solution {

  int[] solution(int[] boardSize, int[] initPosition, int[] initDirection, int k) {
    int n1 = boardSize[0] * 2, n2 = boardSize[1] * 2;
    int l = (((initPosition[0] + initDirection[0] * k) % n1) + n1) % n1, w =
      (((initPosition[1] + initDirection[1] * k) % n2) + n2) % n2;
    return new int[]{l >= boardSize[0] ? boardSize[0] * 2 - l - 1 : l,
      w >= boardSize[1] ? boardSize[1] * 2 - w - 1 : w};
  }

}
