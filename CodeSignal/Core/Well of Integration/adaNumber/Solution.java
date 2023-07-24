public class Solution {

  boolean solution(String line) {
    String digits = "0123456789ABCDEF";
    line = line.replace("_", "").toUpperCase();
    if (!line.matches("[0-9A-F#]+")) {
      return false;
    }
    String[] groups = line.split("#");
    if (groups.length == 1) {
      return line.matches("[0-9]+");
    } else if (line.matches("[0-9]{1,2}#[0-9A-F]+#")) {
      int base = Integer.parseInt(groups[0]);
      System.out.println("base=" + base);
      if (base < 2 || base > 16) {
        return false;
      }
      return groups[1].matches("[" + digits.substring(0, base) + "]+");
    }
    return false;
  }
}
