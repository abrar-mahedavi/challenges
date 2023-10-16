public class Solution {

  int[] solution(String tree) {
    List<Integer> values = new LinkedList<>();
    int maxDepth = 0, depth = 0, i = 0;

    while (i < tree.length()) {
      if (tree.charAt(i) == '(') {
        depth++;
      } else if (tree.charAt(i) == ')') {
        depth--;
      } else if (tree.charAt(i) != ' ') {
        String str = tree.substring(i, tree.indexOf(' ', i));
        i += str.length();

        if (depth > maxDepth) {
          values.clear();
          maxDepth = depth;
        }

        if (depth == maxDepth) {
          values.add(Integer.parseInt(str));
        }
      }

      i++;
    }
    return values.stream().mapToInt(v -> v).toArray();
  }

}
