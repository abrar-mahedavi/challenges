import java.util.*;

class Program {

  public int nonConstructibleChange(int[] coins) {
    // Write your code here.
    Arrays.sort(coins);
    int change=0;
    for(int i= 0;i<coins.length;i++){
      if( coins[i] > change+1){
        return change +1;
      }else {
        change += coins[i] ;
      }
    }
    return change +1;
  }
}
