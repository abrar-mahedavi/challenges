import java.util.HashMap;
import java.util.Map;

public class Solution {

  public static String findSum(final String number1, final String number2) {
    Map<Character, Integer> numbers = new HashMap<Character, Integer>();
    numbers.put('0', 0);
    numbers.put('1', 1);
    numbers.put('2', 2);
    numbers.put('3', 3);
    numbers.put('4', 4);
    numbers.put('5', 5);
    numbers.put('6', 6);
    numbers.put('7', 7);
    numbers.put('8', 8);
    numbers.put('9', 9);

    int maxNumberSize = number1.length() >= number2.length() ? number1.length() : number2.length();
    StringBuilder sb1 = new StringBuilder();
    StringBuilder sb2 = new StringBuilder();
    // For first number added the padding
    if (number1.length() != maxNumberSize) {
      for (int i = 0; i < maxNumberSize - number1.length(); i++) {
        sb1.append("0");
      }
      sb1.append(number1);
      sb2.append(number2);
    }

    // For second number added the padding
    if (number2.length() != maxNumberSize) {
      for (int i = 0; i < maxNumberSize - number2.length(); i++) {
        sb2.append("0");
      }
      sb2.append(number2);
      sb1.append(number1);
    }
    // For first, second number no padding
    if (number1.length() == number2.length()) {
      sb2.append(number2);
      sb1.append(number1);
    }

    // After padding
    // number1 ="010"
    // number2 = "220"
    StringBuilder result = new StringBuilder();
    System.out.println(sb1);
    System.out.println(sb2);

    boolean carryFlag = false;
    for (int i = maxNumberSize - 1; i >= 0; i--) {
      Integer ele1 = numbers.get(sb1.toString().charAt(i));
      Integer ele2 = numbers.get(sb2.toString().charAt(i));
      int sum = 0;
      if (!carryFlag) {
        sum = ele1 + ele2;
      } else {
        sum = ele1 + ele2 + 1;
      }

      if (sum / 10 > 0) {
        carryFlag = true;
      } else {
        carryFlag = false;
      }
      result.append(sum % 10);
    }
    if (carryFlag) {
      result.append("1");
    }
    return result.reverse();

  }

  public static void main(String[] args) {
    String number1 = "10";  // 999
    String number2 = "220"; //999

    System.out.println(findSum(number1, number2);


  }

}


