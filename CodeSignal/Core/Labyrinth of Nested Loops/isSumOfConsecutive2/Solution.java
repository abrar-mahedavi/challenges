public class Solution {

  int solution(int n) {
    int count=0, i, m;
    for(i=1, m=i;m<n;i++,m +=i){
      if((n-m)%(i+1)==0){
        count++;
      }
    }
    return count;
  }
}
