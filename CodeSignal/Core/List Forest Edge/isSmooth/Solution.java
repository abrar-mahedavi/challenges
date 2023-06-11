public class Solution {

  boolean solution(int[] arr) {
    int smooth = (arr.length%2 == 0)?arr[arr.length/2 - 1]+arr[arr.length/2]:arr[arr.length/2];
    return (arr[0] == smooth && smooth == arr[arr.length-1]);
  }

}
