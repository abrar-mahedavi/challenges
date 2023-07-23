`Medium`	`Codewriting` 	`300`

## Description

---

Two arrays are called _similar_ if one can be obtained from another by swapping at most one pair of elements in one of the arrays.

Given two arrays <code>a</code> and <code>b</code>, check whether they are _similar_.

**Example**

- For <code>a = [1, 2, 3]</code> and <code>b = [1, 2, 3]</code>, the output should be
  <code>areSimilar(a, b) = true</code>.
  The arrays are equal, no need to swap any elements.

- For <code>a = [1, 2, 3]</code> and <code>b = [2, 1, 3]</code>, the output should be
  <code>areSimilar(a, b) = true</code>.

  We can obtain <code>b</code> from <code>a</code> by swapping <code>2</code> and <code>1</code> in <code>b</code>.

- For <code>a = [1, 2, 2]</code> and <code>b = [2, 1, 1]</code>, the output should be
  <code>areSimilar(a, b) = false</code>.

  Any swap of any two elements either in <code>a</code> or in <code>b</code> won't make <code>a</code> and <code>b</code> equal.

**Input/Output**

- **[execution time limit] 4 seconds (js)**

- **[input] array.integer a**

  Array of integers.<br>

  _Guaranteed constraints:_<br>
  <code>3 ≤ a.length ≤ 10<sup>5</sup></code>,<br> <code>1 ≤ a[i] ≤ 1000</code>.

- **[input] array.integer b**

  Array of integers of the same length as a.<br>

  _Guaranteed constraints:_<br>
  <code>b.length = a.length</code>,<br> <code>1 ≤ b[i] ≤ 1000h</code>.

- **[output] boolean**
  - <code>true</code> if <code>a</code> and <code>b</code> are similar, <code>false</code> otherwise.
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
