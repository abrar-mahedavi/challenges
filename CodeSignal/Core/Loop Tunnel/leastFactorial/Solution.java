public class Solution {

  int solution(int n) {
    int f = 1;
    for (int i = 1; f < n;i++)
      f *= i;
    return f;
  }

}
