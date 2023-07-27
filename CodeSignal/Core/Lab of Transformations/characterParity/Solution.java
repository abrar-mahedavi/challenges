public class Solution {

  String solution(char symbol) {
    if (symbol < '0' || symbol > '9')
      return "not a digit";
    if ((symbol - '0')%2 == 0)
      return "even";
    return "odd";
  }
}
