`Medium`	`Codewriting` 	`300`

## Description

---

Given a rectangular matrix and integers <code>a</code> and <code>b</code>, consider the union of the <code>a<sup>th</sup></code> row and the <code>b<sup>th</sup></code> (both 0-based) column of the matrix (i.e. all cells that belong either to the a<sup>th</sup> row or to the b<sup>th</sup> column, or to both). Return sum of all elements of that union.

**Example**

For

<code type='preformat'>
matrix = [[1, 1, 1, 1], 
          [2, 2, 2, 2], 
          [3, 3, 3, 3]]
</code>

<code>a = 1</code>, and <code>b = 3</code>, the output should be
<code>crossingSum(matrix, a, b) = 12</code>.

Here <code>(2 + 2 + 2 + 2) + (1 + 3) = 12</code>.

**Input/Output**

- **[execution time limit] 4 seconds (js)**

- **[input] array.array.integer matrix**

  2-dimensional array of integers representing a rectangular matrix.<br>

  _Guaranteed constraints:_<br>
  <code>1 ≤ matrix.length ≤ 100</code>,<br> <code>1 ≤ matrix[0].length ≤ 500</code>,<br> <code>1 ≤ matrix[i][j] ≤ 10<sup>5</sup></code>.

- **[input] integer a**

  A non-negative integer less than the number of matrix rows.<br>

  _Guaranteed constraints:_<br>
  <code>0 ≤ a < matrix.length</code>.

- **[input] integer b**

  A non-negative integer less than the number of matrix columns.<br>

  _Guaranteed constraints:_<br>
  <code>0 ≤ b < matrix[i].length</code>.

* **[output] integer**

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
