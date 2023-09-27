public class Solution {

  int solution(int year, int[] daysOfTheWeek, String[] holidays) {
    int count = 0;
    for (String holiday : holidays) {
      int month = Integer.parseInt(holiday.substring(0,2));
      int day = Integer.parseInt(holiday.substring(3));
      java.time.LocalDate today;
      if (month < 6) {
        today = java.time.LocalDate.of(year + 1, month, day);
      } else {
        today = java.time.LocalDate.of(year, month, day);
      }
      if (IntStream.of(daysOfTheWeek).anyMatch(x -> x == today.getDayOfWeek().getValue())) {
        count++;
      }
    }
    return count;
  }

}
