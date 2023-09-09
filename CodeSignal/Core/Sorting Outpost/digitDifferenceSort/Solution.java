public class Solution {

  int[] solution(int[] a) {
    List<Integer> l = IntStream.range(0, a.length).boxed().collect(Collectors.toList());
    Collections.sort(l, (x, y) ->
      (Arrays.stream(Integer.toString(a[y]).split("")).mapToInt(i -> Integer.parseInt(i)).max()
        .getAsInt() -
        Arrays.stream(Integer.toString(a[y]).split("")).mapToInt(i -> Integer.parseInt(i)).min()
          .getAsInt()) -
        (Arrays.stream(Integer.toString(a[x]).split("")).mapToInt(i -> Integer.parseInt(i)).max()
          .getAsInt() -
          Arrays.stream(Integer.toString(a[x]).split("")).mapToInt(i -> Integer.parseInt(i)).min()
            .getAsInt()));
    Collections.reverse(l);
    return l.stream().mapToInt(i -> a[i]).toArray();
  }


}
