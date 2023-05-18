public class Solution {

  int solution(int[] a) {
    int sum = 0;
    for(int i = 0, e = 0; i < a.length; i++, e += 8)
      sum += a[i] * Math.pow(2, e);
    return sum;
  }

}
