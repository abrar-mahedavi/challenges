public class Solution {

  int solution(int[] a) {
    long sum = 0;
    for (int k : a) sum += k;
    long third = sum / 3, numOfCuts = 0, currSum = 0, ans = 0;
    for (int i = 0; i < a.length - 1; i++){
      currSum += a[i];
      if (currSum == 2*third)
        ans += numOfCuts;
      if (currSum == third)
        numOfCuts++;
    }
    return (int) ans;
  }


}
