public class Solution {

  boolean solution(String p) {
    return (~p.hashCode() & 1) == 0;
  }


}
