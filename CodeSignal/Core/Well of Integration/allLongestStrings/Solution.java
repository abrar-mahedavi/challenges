public class Solution {

  String[] solution(String[] inputArray) {
    int mL = 0;
    List<String> l = new ArrayList<>();
    for (String str : inputArray) {
      if (str.length() > mL) {
        mL = str.length();
        l.clear();
      }
      if (str.length() == mL)
        l.add(str);
    }
    String[] ss = new String[l.size()];
    for (int i=0; i<l.size(); i++) {
      ss[i] = l.get(i);
    }
    return ss;
  }

}
