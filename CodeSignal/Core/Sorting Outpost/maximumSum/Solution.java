public class Solution {

  int solution(int[] a, int[][] q) {
    int[] sumCount = new int[a.length];
    for (int i = 0; i < q.length; i++) {
      for (int j = q[i][0]; j <= q[i][1]; j++) {
        sumCount[j]++;
      }
    }
    Arrays.sort(a);
    Arrays.sort(sumCount);
    int answer = 0;
    for (int i = 0; i < a.length; i++) {
      answer += a[i] * sumCount[i];
    }

    return answer;
  }


}
