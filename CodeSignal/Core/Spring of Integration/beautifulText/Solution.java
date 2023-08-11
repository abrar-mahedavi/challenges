public class Solution {

  boolean solution(String inputString, int l, int r) {
    while (l <= r) {
      int i = l++, x = inputString.length();
      for (; i < x; i += l)
        if (inputString.charAt(i) > 32)
          break;

      if (i == x)
        return true;
    }

    return false;
  }
}
