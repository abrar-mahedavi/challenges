public class Solution {

  int solution(int[] a) {
    List<Integer> l = Arrays.stream(a)
      .map(i -> Arrays.stream(Integer.toString(i).split("")).mapToInt(j ->
        Integer.parseInt(j)).reduce((x, y) -> x * y).getAsInt()).boxed()
      .collect(Collectors.toList());
    return new HashSet<Integer>(l).size();
  }

}
