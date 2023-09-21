public class Solution {

  boolean solution(String time) {
    String [] strarr = time.split(":");
    int hours = Integer.parseInt(strarr[0]);
    int minutes = Integer.parseInt(strarr[1]);
    return (hours < 24 && minutes < 60);
  }

}
