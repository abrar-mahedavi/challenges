public class Solution {

  boolean solution(int[] a, int[] b) {
    int[] di = IntStream.range(0, a.length).filter(i -> a[i] != b[i]).toArray();
    return di.length == 0 || (di.length == 2 && a[di[0]] == b[di[1]] && a[di[1]] == b[di[0]]);
  }


}
