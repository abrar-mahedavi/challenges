public class Solution {

  int[] solution(int[] a) {
    int counter = 0;
    for (int i=a.length-1; i>=0; i--) {
      if (a[i] == 1)
        counter++;
      a[i] = (a[i]+counter)%2;
    }
    return a;
  }

}
