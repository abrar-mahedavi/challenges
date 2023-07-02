public class Solution {

  boolean solution(String filename1, String filename2) {
    return filename1.compareTo(filename2) * filename1.compareToIgnoreCase(filename2) < 0;
  }

}
