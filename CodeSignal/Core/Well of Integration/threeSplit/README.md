`Medium`	`Codewriting` 	`300`

## Description

---

You have a long strip of paper with integers written on it in a single line from left to right. You wish to cut the paper into exactly three pieces such that each piece contains at least one integer and the sum of the integers in each piece is the same. You cannot cut through a number, i.e. each initial number will unambiguously belong to one of the pieces after cutting. How many ways can you do it?

It is guaranteed that the sum of all elements in the array is divisible by <code>3</code>.

**Example**

For <code>a = [0, -1, 0, -1, 0, -1]</code>, the output should be
<code>threeSplit(a) = 4</code>.

Here are all possible ways:

- <code>[0, -1][0, -1] [0, -1]</code>
- <code>[0, -1][0, -1, 0] [-1]</code>
- <code>[0, -1, 0][-1, 0] [-1]</code>
- <code>[0, -1, 0][-1] [0, -1]</code>

**Input/Output**

- **[execution time limit] 4 seconds (js)**

- **[input] array.integer a**

  _Guaranteed constraints:_<br>
  <code>5 ≤ a.length ≤ 10<sup>4</sup></code>,<br> <code>-10<sup>8</sup> ≤ a[i] ≤ 10<sup>8</sup></code>.

- **[output] integer**
  - It's guaranteed that for the given test cases the answer always fits signed <code>32</code>-bit integer type.
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
