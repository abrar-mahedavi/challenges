public class Solution {

  int solution(String letter) {
    Pattern pattern = Pattern.compile("([a-z]+)[^a-z]+\\1(?:$|[^a-z])",Pattern.CASE_INSENSITIVE);
    Matcher matcher = pattern.matcher(letter);

    int res = 0;
    while (matcher.find()) {
      res++;
    }
    return res;
  }

}
