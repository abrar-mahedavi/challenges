`Medium`	`Codewriting` 	`300`


## Description

---

Define an integer's roundness as the number of trailing zeroes in it.

Given an integer <code>n</code>, check if it's possible to increase <code>n</code>'s roundness by swapping some pair of its digits.

**Example**

- For <code>n = 902200100</code>, the output should be
  <code>increaseNumberRoundness(n) = true</code>.

  One of the possible ways to increase roundness of <code>n</code> is to swap digit <code>1</code> with digit <code>0</code> preceding it: roundness of <code>902201000</code> is <code>3</code>, and roundness of <code>n</code> is <code>2</code>.

  For instance, one may swap the leftmost <code>0</code> with <code>1</code>.

- For <code>n = 11000</code>, the output should be
  <code>increaseNumberRoundness(n) = false</code>.

  _Roundness_ of <code>n</code> is <code>3</code>, and there is no way to increase it.

**Input/Output**

- **[execution time limit] 4 seconds (js)**

- **[input] integer n**

  A positive integer.

  _Guaranteed constraints:_<br>
  <code type='math/tex'>100 \leq n \leq 10^9</code>.

- **[output] boolean**
  - <code>true</code> if it's possible to increase <code>n</code>'s roundness, <code>false</code> otherwise.

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
