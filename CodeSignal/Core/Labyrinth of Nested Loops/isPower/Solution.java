public class Solution {

  boolean solution(int n) {
    for( int i = 2 ; i < 5 ; i++ ){
      if(Math.round(Math.pow(n , 1f/i) * 1000f) / 1000f % 1 == 0){
        return true;
      };
    }
    return false;
  }

}
