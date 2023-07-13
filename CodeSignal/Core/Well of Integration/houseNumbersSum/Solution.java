public class Solution {

  int solution(int[] inputArray) {
    int i = 0;
    int sum = 0;
    while(inputArray[i] != 0)sum += inputArray[i++];

    return sum;
  }

}
