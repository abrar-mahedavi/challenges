public class Solution {

  int[] solution(int[] arr) {
    if (arr.length%2 == 1)
      return arr;
    int[] newArr = new int[arr.length-1];
    for (int i = 0 ; i < arr.length/2; i++) {
      newArr[i] += arr[i];
      newArr[newArr.length - 1 - i] += arr[arr.length - 1 - i];
    }
    return newArr;
  }

}
