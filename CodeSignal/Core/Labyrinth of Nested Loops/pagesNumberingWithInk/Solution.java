public class Solution {

  int solution(int current, int numberOfDigits) {
    int currentDigits=0;
    int n = current;
    while (n>0)
    {
      n=n/10;
      currentDigits++;
      numberOfDigits--;
    }
    current++;

    while (currentDigits<=numberOfDigits)
    {
      n = current;
      currentDigits=0;
      while (n>0)
      {
        n=n/10;
        currentDigits++;
      }
      current++;
      numberOfDigits-=currentDigits;
    }
    return current-1;
  }


}
