public class Solution {

  public static int findMaxElement(int[] numbers) {
    int maximumElement = numbers[0];
    for(int i = 1;i<numbers.length;i++){
      if (numbers[i]> maximumElement)
        maximumElement = numbers[i];
    }
    return maximumElement;
  }

  public static void main(String[] args) {
    System.out.println(findMaxElement(new int[]{2,4,8,-1,4,7,0}));
    System.out.println(findMaxElement(new int[]{-2,-4,-8,-1,-4,-7}));
  }
}
