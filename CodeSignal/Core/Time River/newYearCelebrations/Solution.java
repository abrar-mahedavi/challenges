public class Solution {

  int solution(String takeOffTime, int[] minutes) {
    int h1 = Integer.parseInt(takeOffTime.substring(0, 2));
    int hours = h1 != 0 ? h1 : 24;
    int min = Integer.parseInt(takeOffTime.substring(3));
    int timeInMin = hours * 60 + min;
    int dayInMin = 60 * 24;
    int result = 0;
    int i = 0;
    while (i < minutes.length) {
      int stopTime = timeInMin + (i == 0 ? minutes[i] : minutes[i] - minutes[i - 1]);
      if (timeInMin <= dayInMin && stopTime >= dayInMin) {
        result++;
      }
      timeInMin = stopTime - 60;
      i++;
    }
    return timeInMin <= dayInMin ? result + 1 : result;
  }

}
