`Medium`	`Codewriting` 	`300`

## Description

---

The _longest diagonals_ of a square matrix are defined as follows:

- the first _longest diagonal_ goes from the top left corner to the bottom right one;
- the second _longest diagonal_ goes from the top right corner to the bottom left one.

Given a square `matrix`, your task is to swap its _longest diagonals_ by exchanging their elements at the corresponding positions.

**Example**

For

<code type='preformat'>
matrix = [[1, 2, 3],
          [4, 5, 6],
          [7, 8, 9]]
</code>

the output should be

<code type='preformat'>
swapDiagonals(matrix) = [[3, 2, 1],
                         [4, 5, 6],
                         [9, 8, 7]]
</code>

**Input/Output**

- **[execution time limit] 4 seconds (js)**

- **[input] array.array.integer matrix**

  _Guaranteed constraints:_<br>
  <code>1 ≤ matrix.length ≤ 100</code>,<br> <code>matrix.length = matrix[i].length</code>,<br> <code>1 ≤ matrix[i][j] ≤ 10<sup>5</sup></code>.

* **[output] array.array.integer**
  - Matrix with swapped diagonals.

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
