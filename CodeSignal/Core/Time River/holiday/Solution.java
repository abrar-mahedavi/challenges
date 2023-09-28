public class Solution {

  int solution(int x, String weekDay, String month, int yearNumber) {
    java.time.Month mon = java.time.Month.valueOf(month.toUpperCase());
    java.time.DayOfWeek dayOfWeek = java.time.DayOfWeek.valueOf(weekDay.toUpperCase());

    java.time.LocalDate d = java.time.LocalDate.of(yearNumber, mon.getValue(), 1);

    while (d.getDayOfWeek() != dayOfWeek) {
      d = d.plusDays(1);
    }
    d = d.plusWeeks(x - 1);
    if (d.getMonth() != mon) {
      return -1;
    }
    return d.getDayOfMonth();
  }


}
