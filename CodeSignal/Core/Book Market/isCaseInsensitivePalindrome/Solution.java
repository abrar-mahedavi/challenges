public class Solution {

  boolean solution(String inputString) {
    return new StringBuilder(inputString).reverse().toString().equalsIgnoreCase(inputString);
  }

}
