public class Solution {

  String solution(String currMonth) {
    int month = Integer.parseInt(currMonth.substring(0, 2));
    int year = Integer.parseInt(currMonth.substring(3));

    java.time.LocalDate today = java.time.LocalDate.of(year, month, 1);

    boolean done = false;
    while (!done) {
      today = today.plusMonths(1);
      if (today.getDayOfWeek() == java.time.DayOfWeek.MONDAY) {
        done = true;
      }
    }
    String result = today.getMonthValue() + "-" + today.getYear();
    if (today.getMonthValue() < 10) {
      result = "0" + result;
    }
    return result;
  }


}
