`Medium`	`Codewriting` 	`300`

## Description

---

Given a string, check if it can become a palindrome (a palindrome is a string that reads the same left-to-right and right-to-left) through a case change of some (possibly, none) letters.

**Example**

- For <code>inputString = "AaBaa"</code>, the output should be
  <code>isCaseInsensitivePalindrome(inputString) = true</code>.

<code>"aabaa"</code> is a palindrome as well as <code>"AABAA"</code>, <code>"aaBaa"</code>, etc.

- For <code>inputString = "abac"</code>, the output should be
  <code>isCaseInsensitivePalindrome(inputString) = false</code>.

All the strings which can be obtained via changing case of some group of letters, i.e. <code>"abac"</code>, <code>"Abac"</code>, <code>"aBAc"</code> and 13 more, are not palindromes.

**Input/Output**

- **[execution time limit] 4 seconds (js)**

- **[input] string inputString**

  Non-empty string consisting of English letters.<br>

  _Guaranteed constraints:_<br>
  <code>4 ≤ inputString.length ≤ 10</code>.

- **[output] boolean**

## [Java] Syntax Tips

``` java
// Prints help message to the console
// Returns a string
// 
// Globals declared here will cause a compilation error,
// declare variables inside the function instead!
String helloWorld(String name) {
    System.out.println("This prints to the console when you Run Tests");
    return "Hello, " + name;
}
```
