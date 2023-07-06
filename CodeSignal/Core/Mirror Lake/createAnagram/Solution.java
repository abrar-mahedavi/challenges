public class Solution {

  int solution(String s, String t) {
    int[] count = new int[256];
    for (byte b : s.getBytes())
      count[b]++;
    for (byte b : t.getBytes())
      count[b]--;
    int total = 0;
    for (int i : count)
      if (i > 0)
        total += i;
    return total;
  }

}
