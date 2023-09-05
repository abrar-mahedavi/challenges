public class Solution {

  String[] solution(String[] inputArray) {
    Arrays.sort(inputArray, (x, y) -> x.length() - y.length());
    return inputArray;
  }

}
