public class Solution {

  String[] solution(String[] names) {
    Set<String> files = new HashSet<>();
    return Arrays.stream(names).map(n -> {
      if (!files.contains(n)) return n;
      else {
        for (int i = 1; ; i++)
          if (!files.contains(n + "(" + i + ")"))
            return n + "(" + i + ")";
      }
    }).peek(files::add).toArray(String[]::new);
  }

}
