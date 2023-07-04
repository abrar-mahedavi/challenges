`Medium`	`Codewriting` 	`300`

## Description

---

Given two strings <code>a</code> and <code>b</code>, both consisting only of lowercase English letters, your task is to calculate how many strings equal to <code>a</code> can be constructed using only letters from the string <code>b</code>? Each letter can be used only once and in one string only.

**Example**

- For <code>a = "abc"</code> and <code>b = "abccba"</code>, the output should be <code>stringsConstruction(a, b) = 2</code>.

  We can construct <code>2</code> strings <code>a = "abc"</code> using only letters from the string <code>b</code>.

- For <code>a = "ab"</code> and <code>b = "abcbcb"</code>, the output should be <code>stringsConstruction(a, b) = 1</code>.

**Input/Output**

- **[execution time limit] 4 seconds (js)**

- **[input] string a**

  String to construct, containing only lowercase English letters.<br>

  _Guaranteed constraints:_<br>
  <code>1 ≤ a.length ≤ 10<sup>5</sup></code>.

- **[input] string b**

  String containing needed letters, containing only lowercase English letters.<br>

  _Guaranteed constraints:_<br>
  <code>1 ≤ b.length ≤ 10<sup>5</sup></code>.

- **[output] integer**
  - The number of strings <code>a</code> that can be constructed from the string <code>b</code>.

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
