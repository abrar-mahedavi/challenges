`Medium`	`Codewriting` 	`300`

## Description

---

A noob programmer was given two simple tasks: sum and sort the elements of the given array
<code>a = [a<sub>1</sub>, a<sub>2</sub>, ..., a<sub>n</sub>]</code>. He started with summing and did it easily, but decided to store the sum he found in some random position of the original array which was a bad idea. Now he needs to cope with the second task, sorting the original array a, and it's giving him trouble since he modified it.

Given the array <code>shuffled</code>, consisting of elements <code>a<sub>1</sub>, a<sub>2</sub>, ..., a<sub>n</sub>, a<sub>1</sub> + a<sub>2</sub> + ... + a<sub>n</sub></code> in random order, return the sorted array of original elements <code>a<sub>1</sub>, a<sub>2</sub>, ..., a<sub>n</sub></code>.

**Example**

- For <code>shuffled = [1, 12, 3, 6, 2]</code>, the output should be<br>
  <code>shuffledArray(shuffled) = [1, 2, 3, 6]</code>.

  <code>1 + 3 + 6 + 2 = 12</code>, which means that <code>1, 3, 6</code> and <code>2</code> are original elements of the array.

- For <code>shuffled = [1, -3, -5, 7, 2]</code>, the output should be<br>
  <code>shuffledArray(shuffled) = [-5, -3, 2, 7]</code>.

**Input/Output**

- **[execution time limit] 4 seconds (js)**

- **[input] array.integer shuffled**

  Array of at least two integers. It is guaranteed that there is an index <code>i</code> such that <code>shuffled[i] = shuffled[0] + ... + shuffled[i - 1] + shuffled[i + 1] + ... + shuffled[n]</code>.

  _Guaranteed constraints:_<br>
  <code>2 ≤ shuffled.length ≤ 10<sup>4</sup></code>,<br>
  <code>-5 · 10<sup>4</sup> ≤ shuffled[i] ≤ 5 · 10<sup>4</sup></code>.

* **[output] array.integer**
  - A sorted array of <code>shuffled.length - 1</code> elements

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
