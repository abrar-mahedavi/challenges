public class Solution {

  String solution(String table, int row, int column) {
    String[] lines = table.split("<");

    int trCount = 0;
    for (int i = 0; i < lines.length; i++) {
      String line = lines[i];
      if (line.equals("tr>")) {
        trCount++;
        if (trCount != row+1) {continue;}
        if (lines[i + 1].startsWith("th")) {
          continue;
        }
        i++;
        int tdCount = 0;
        String value = "";
        while (tdCount <= column) {
          line = lines[i];
          value = line.substring(3);
          if (value.equals(">")) {
            return "No such cell";
          }
          i += 2;
          tdCount++;
        }
        return value;
      }
    }
    return "No such cell";
  }


}
