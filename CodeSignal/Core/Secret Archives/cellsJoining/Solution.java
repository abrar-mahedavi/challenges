public class Solution {

  String[] solution(String[] table, int[][] coords) {
    int row = -1;
    for (int i = 0; i < table.length; i++) {
      if (table[i].startsWith("+")) {
        row++;
      }
      Pattern p = Pattern.compile("(?<p1>(?<sep>.)(.+?\\k<sep>){" + coords[0][1] + "})"
        + "(?<p2>.+?(\\k<sep>.+?){" + (coords[1][1] - coords[0][1]) + "}(?=\\k<sep>))"
        + "(?<p3>.*)");
      String rowSepReplace = "$0";
      String rowColReplace = "$0";
      String singlePlusReplace = "$0";
      if (row == coords[1][0] && table[i].startsWith("+")) {
        if (i == 0) {
          rowSepReplace = "-";
        }
      } else if (row == coords[0][0] + 1 && table[i].startsWith("+")) {
        if (i == table.length - 1) {
          // last row
          rowSepReplace = "-";
        }
      } else if (row >= coords[1][0] && row <= coords[0][0]) {
        rowSepReplace = " ";
        if (table[i].startsWith("+")) {
          rowColReplace = " ";
          singlePlusReplace = "|";
        }
      }
      if (!rowSepReplace.equals("$0") || !rowColReplace.equals("$0")) {
        Matcher m = p.matcher(table[i]);
        if (!m.matches()) throw new AssertionError(table[i] + " ~= " + p);
        table[i] = m.group("p1").replaceAll("^\\+$", singlePlusReplace) +
          m.group("p2")
            .replaceAll("[^\\" + m.group("sep") + "]", rowColReplace)
            .replaceAll("\\" + m.group("sep"), rowSepReplace)
          + m.group("p3").replaceAll("^\\+$", singlePlusReplace);
      }
    }


    return table;
  }

}
