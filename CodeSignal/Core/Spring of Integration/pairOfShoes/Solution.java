public class Solution {

  boolean solution(int[][] shoes) {
    int[] types = new int[101];
    for (int[] i : shoes) {
      if (i[0] == 0)
        types[i[1]]--;
      else
        types[i[1]]++;
    }
    for (int i : types)
      if (i != 0)
        return false;
    return true;
  }
}
