public class Solution {

  int solution(int n) {
    return (n / 60) % 10 + (n / 60 / 10) % 10 + (n % 60) % 10 + ((n % 60) / 10) % 10;
  }

}
