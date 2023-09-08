`Medium`	`Codewriting` 	`300`

## Description

---

Given a rectangular matrix of integers, check if it is possible to rearrange its rows in such a way that all its columns become strictly increasing sequences (read from top to bottom).

**Example**

- For

  <code type='preformat'>
    matrix = [[2, 7, 1], 
              [0, 2, 0], 
              [1, 3, 1]]
  </code>

  the output should be

  <code type='preformat'>
    rowsRearranging(matrix) = false;
  </code>

- For

  <code type='preformat'>
    matrix = [[6, 4], 
              [2, 2], 
              [4, 3]]
  </code>

  the output should be

  <code type='preformat'>
    rowsRearranging(matrix) = true.
  </code>

**Input/Output**

- **[execution time limit] 4 seconds (js)**

- **[input] array.array.integer matrix**

  A 2-dimensional array of integers.

  _Guaranteed constraints:_<br>
  <code>1 ≤ matrix.length ≤ 10</code>,<br>
  <code>1 ≤ matrix[0].length ≤ 10</code>,<br>
  <code>-300 ≤ matrix[i][j] ≤ 300</code>.

* **[output] boolean**

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
