public class Solution {

  int solution(String[] inputArray, String result) {
    int count = 0;
    for (int i = 0; i < inputArray.length; i++) {
      for (int j = i + 1; j < inputArray.length; j++) {
        if (possibleResult(inputArray[i], inputArray[j], result)) {
          count++;
        }
      }
    }
    return count;
  }

  boolean possibleResult(String s1, String s2, String result) {
    boolean possible = true;
    for (int i = 0; i < s1.length(); i++) {
      if (!(s1.charAt(i) == result.charAt(i) || s2.charAt(i) == result.charAt(i))) {
        possible = false;
      }
    }
    return possible;
  }

}
