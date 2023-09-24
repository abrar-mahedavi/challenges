public class Solution {

  String solution(String someTime, String leavingTime) {
    SimpleDateFormat fomat = new SimpleDateFormat("yyyy-MM-dd HH:mm");
    try {
      Date d1 = fomat.parse(someTime);
      Date d2 = fomat.parse(leavingTime);
      long period = d2.getTime() - d1.getTime();
      long timed3 = d1.getTime() - period;
      Calendar cal = Calendar.getInstance();
      cal.setTimeInMillis(timed3);
      Date d3 = cal.getTime();
      return fomat.format(d3);
    } catch (ParseException e) {
      e.printStackTrace();
      return "";
    }
  }

}
