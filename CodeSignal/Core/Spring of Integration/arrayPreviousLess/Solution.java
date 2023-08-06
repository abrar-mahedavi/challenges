public class Solution {

  int[] solution(int[] items) {
    for(int i = items.length-1; i >= 0; i--){
      for(int j = i; j >= 0; j--){
        if(items[j] < items[i]){
          items[i] = items[j];
          break;
        }
        if(j == 0)
          items[i] = -1;
      }
    }
    return items;
  }
}
