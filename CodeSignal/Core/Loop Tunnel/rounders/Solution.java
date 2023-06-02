public class Solution {

  int solution(int value) {
    int p = 1;
    while(value > 10) {
      p *= 10;
      value = (value/10) + ((value%10<5)?0:1);
    }
    return value*p;
  }

}
