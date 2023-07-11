public class Solution {

  int solution(int[] divisors, int k) {
    Set<Integer> friendMap = new HashSet<Integer>();
    for (int i=1; i<=k; i++) {
      int signature = 0;
      for (int j=0; j<divisors.length; j++) {
        if (i%divisors[j]==0) {
          signature += (int)Math.pow(2,j);
        }
      }
      friendMap.add(signature);
    }
    return friendMap.size();
  }


}
