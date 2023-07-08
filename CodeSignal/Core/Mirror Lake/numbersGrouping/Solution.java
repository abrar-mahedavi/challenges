public class Solution {

  int solution(int[] a) {
    Set<Integer> groups = new HashSet<>();
    for (int num : a) {
      groups.add((num - 1) / 10000);
    }
    return a.length + groups.size();
  }

}
