`Medium`	`Codewriting` 	`300`

## Description

---

Given a rectangular matrix containing only digits, calculate the number of different <code>2 × 2</code> squares in it.

**Example**

For

<code>
matrix = [[1, 2, 1],
          [2, 2, 2],
          [2, 2, 2],
          [1, 2, 3],
          [2, 2, 1]]
</code>

the output should be
<code>differentSquares(matrix) = 6</code>.

Here are all <code>6</code> different <code>2 × 2</code> squares:

- 1 2<br>
  2 2
- 2 1<br>
  2 2
- 2 2<br>
  2 2
- 2 2<br>
  1 2
- 2 2<br>
  2 3
- 2 3<br>
  2 1

**Input/Output**

- **[execution time limit] 4 seconds (js)**

- **[input] array.array.integer matrix**

  _Guaranteed constraints:_<br>
  <code>1 ≤ matrix.length ≤ 100</code>,<br>
  <code>1 ≤ matrix[i].length ≤ 100</code>,<br>
  <code>0 ≤ matrix[i][j] ≤ 9</code>.

- **[output] integer**
  - The number of different <code>2 × 2</code> squares in <code>matrix</code>.

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
