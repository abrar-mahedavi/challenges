public class Solution {

  int solution(int[] inputArray) {
    int n = inputArray.length, i = 1;

    while (n>1) {
      for (int j = 0; j < n/2; j ++) {
        inputArray[j] = i==1?inputArray[2*j]+ inputArray[2*j+1]:inputArray[2*j]* inputArray[2*j+1];
      }
      n /= 2;
      i = 1-i;
    }

    return inputArray[0];
  }

}
