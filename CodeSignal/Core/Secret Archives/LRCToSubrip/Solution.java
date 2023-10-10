public class Solution {

  String[] solution(String[] lrcLyrics, String songLength) {
    int len = lrcLyrics.length;
    String[] SubRip = new String[4 * len - 1];
    String prevTime = "";
    for (int i = 0; i < len; i++) {
      String line = lrcLyrics[i];
      SubRip[4 * i] = Integer.toString(i + 1);

      int index = line.indexOf("]");
      String time = line.substring(1, index);

      int minute = Integer.parseInt(time.substring(0, 2));
      int hour = minute / 60;
      minute = minute % 60;

      String seconds = time.substring(3);
      seconds = seconds.replace(".", ",") + "0";
      time = String.format("%02d:%02d:%s", hour, minute, seconds);
      if (i > 0) {
        SubRip[4 * (i - 1) + 1] = prevTime + " --> " + time;
      }
      prevTime = time;

      if (line.length() >= index + 2) {
        SubRip[4 * i + 2] = line.substring(index + 2);
      } else {
        SubRip[4 * i + 2] = "";
      }

      if (i < lrcLyrics.length - 1) {
        SubRip[4 * i + 3] = "";
      }
    }
    SubRip[4 * (len - 1) + 1] = prevTime + " --> " + songLength + ",000";
    return SubRip;
  }

}
