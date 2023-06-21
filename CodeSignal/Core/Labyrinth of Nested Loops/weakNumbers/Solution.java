public class Solution {

  int[] solution(int n) {
    int[] result = new int[2];
    int[] d = new int[n];
    int[] m = new int[n];
    for(int i = 1; i<=n; i++){
      int div = 0;
      for(int x = 1; x<=i;x++){
        if(i%x == 0){
          div++;
        }
      }
      d[i-1] = div;
    }
    int max = 0;
    for(int i = n-1; i>=0; i--){
      int count = 0;
      for(int x = i-1; x>=0; x--){
        if(d[i]<d[x]){
          count++;
        }
      }
      if(count>max)max = count;
      m[i] = count;
    }
    result[0] = max;
    for(int i = 0; i<n; i++){
      if(m[i] == max){
        result[1]++;
      }
    }
    return result;
  }

}
