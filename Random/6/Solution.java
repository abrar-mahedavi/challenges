public class Solution {

  public static void findOperations(List<Integer> source,List<Integer> destination) {
    for(int i=0;i< source.size();i++){
      final int finalI = i;
      final List<Integer> collect = destination.stream().filter(value -> value == source.get(finalI)).collect(Collectors.toList());
      if(collect.size() == 0){
        System.out.println("ADD =>"+source.get(i));
      }
    }

    for(int i=0;i< destination.size();i++){
      final int finalI = i;
      final List<Integer> collect = source.stream().filter(value -> value == destination.get(finalI)).collect(Collectors.toList());
      if(collect.size() == 0){
        System.out.println("DELETE =>"+destination.get(i));
      }
    }
  }

  public static void main(String[] args) {
    List<Integer> revenue = new ArrayList<>(Arrays.asList(1,2,3));
    List<Integer> one = new ArrayList<>(Arrays.asList(1,4));
    findOperations(revenue,one);
  }
}
