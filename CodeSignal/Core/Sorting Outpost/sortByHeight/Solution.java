public class Solution {

  int[] solution(int[] a) {
    PrimitiveIterator.OfInt it = Arrays.stream(a).filter(n -> n != -1).sorted().iterator();
    return Arrays.stream(a).map(n -> n == -1 ? n : it.next()).toArray();
  }

}
