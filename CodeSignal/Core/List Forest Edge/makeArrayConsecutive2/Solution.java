public class Solution {

  int solution(int[] statues) {
    Arrays.sort(statues);
    return statues[statues.length-1] - statues[0] - statues.length + 1;
  }

}
