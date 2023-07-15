public class Solution {

  int[] solution(int legs) {
    int[] human = new int[legs / 4 + 1];
    int index = legs / 4;
    while (legs >= 0) {
      human[index] = legs / 2;
      legs -= 4;
      index--;
    }
    return human;
  }


}
