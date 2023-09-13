`Medium`	`Codewriting` 	`300`

## Description

---

Given a position of a knight on the standard chessboard, find the number of different moves the knight can perform.

The knight can move to a square that is two squares horizontally and one square vertically, or two squares vertically and one square horizontally away from it. The complete move therefore looks like the letter L. Check out the image below to see all valid moves for a knight piece that is placed on one of the central squares.

![](./img_2.png)

**Example**

- For <code>cell = "a1"</code>, the output should be
  <code>chessKnightMoves(cell) = 2</code>.

  ![](./img.png)

- For <code>cell = "c2"</code>, the output should be
  <code>chessKnightMoves(cell) = 6</code>.

  ![](./img_1.png)

**Input/Output**

- **[execution time limit] 4 seconds (js)**

- **[input] string cell**

  String consisting of 2 letters - coordinates of the knight on an 8 × 8 chessboard in chess notation.<br>

  _Guaranteed constraints:_<br>
  <code>cell.length = 2</code>,<br>
  <code>'a' ≤ cell[0] ≤ 'h'</code>,<br>
  <code>1 ≤ cell[1] ≤ 8</code>.

  Note (Chess notation): Each square of the chessboard is identified by a unique coordinate pair—a letter and a number. The vertical columns of squares from white's left to the right are labeled 'a' through 'h'. The horizontal rows of squares are numbered 1 to 8 starting from white's side of the board. Thus each square has a unique identification as a string consisting of two characters: the first is the column label, and the second in the row number.

  ![](./img_3.png)

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
