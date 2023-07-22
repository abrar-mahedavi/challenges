`Medium`	`Codewriting` 	`300`

## Description

---

Given a positive integer number and a certain length, we need to modify the given number to have a specified length. We are allowed to do that either by cutting out leading digits (if the number needs to be shortened) or by adding <code>0s</code> in front of the original number.

**Example**

- For <code>number = 1234</code> and <code>width = 2</code>, the output should be
  <code>integerToStringOfFixedWidth(number, width) = "34"</code>;
- For <code>number = 1234</code> and <code>width = 4</code>, the output should be
  <code>integerToStringOfFixedWidth(number, width) = "1234"</code>;
- For <code>number = 1234</code> and <code>width = 5</code>, the output should be
  <code>integerToStringOfFixedWidth(number, width) = "01234"</code>.

**Input/Output**

- **[execution time limit] 4 seconds (js)**

- **[input] integer number**

  A non-negative integer.<br>

  _Guaranteed constraints:_<br>
  <code>0 ≤ number ≤ 10<sup>9</sup></code>.

- **[input] integer width**

  A positive integer representing the desired length.<br>

  _Guaranteed constraints:_<br>
  <code>1 ≤ width ≤ 50</code>.

- **[output] string**
  - The modified version of <code>number</code> as described above.

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
