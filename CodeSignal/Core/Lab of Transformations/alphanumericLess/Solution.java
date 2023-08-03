public class Solution {

  boolean solution(String s1, String s2) {
    int index1 = 0;
    int index2 = 0;

    int zeroCount1 = 0;
    int zeroCount2 = 0;

    while (index1 < s1.length() && index2 < s2.length()) {
      char c1 = s1.charAt(index1);
      char c2 = s2.charAt(index2);

      if (isLetter(c1) && isLetter(c2)) {
        // compare letters
        if (c1 != c2) {
          return c1 < c2;
        }
        index1++;
        index2++;
      } else if (!isLetter(c1) && isLetter(c2)) {
        return true;
      } else if (isLetter(c1) && !isLetter(c2)) {
        return false;
      } else {
        // both numbers
        String[] n1 = getNum(index1, s1);
        String[] n2 = getNum(index2, s2);

        if (n1[1].length() != n2[1].length()) {
          System.out.println(n1[1] + " " + n2[1]);
          return n1[1].length() < n2[1].length();
        }
        int comp = n1[1].compareTo(n2[1]);
        if (comp != 0) {
          return comp < 0;
        }

        int z1 = Integer.parseInt(n1[0]);
        int z2 = Integer.parseInt(n2[0]);

        if (zeroCount1 == 0 && zeroCount2 ==0 && z1 != z2) {
          zeroCount1 = z1;
          zeroCount2 = z2;
        }
        index1 = Integer.parseInt(n1[2]);
        index2 = Integer.parseInt(n2[2]);

      }
    }

    if (index1 < s1.length()) {
      return false;
    } else if (index2 < s2.length()) {
      return true;
    }

    return zeroCount1 > zeroCount2;
  }

  boolean isLetter (char c) {
    return c >= 'a' && c <= 'z';
  }

  String[] getNum(int index, String str) {
    boolean leading = true;
    int zeroes = 0;
    String[] result = new String[3];

    String numStr = "";

    while (index < str.length()) {
      char c = str.charAt(index);
      if (isLetter(c)) {
        break;
      }
      if (c == '0' && leading) {
        zeroes++;
      } else {
        leading = false;
        numStr = numStr + c;
      }
      index++;
    }

    result[0] = Integer.toString(zeroes);
    result[1] = numStr;
    result[2] = Integer.toString(index);

    return result;
  }
}
