public class Solution {

  public static void palindrome(String word) {
    int i = 0;
    int j = word..length() - 1;
    while (i <= j) {
      if (word.charAt(i) != word.charAt(j)) return false;
      i++;
      j--;
    }
    return true;
  }

  public static void main(String[] args) {
    palindrome("123321");
    palindrome("dad");
    palindrome("amanaplanacanalpanama");
  }
}
