`Medium`	`Codewriting` 	`300`

## Description

---

Given the positions of a white <code>bishop</code> and a black <code>pawn</code> on the standard chess board, determine whether the bishop can capture the pawn in one move.

The bishop has no restrictions in distance for each move, but is limited to diagonal movement. Check out the example below to see how it can move:

![](./img_2.png)

**Example**

- For <code>bishop = "a1"</code> and <code>pawn = "c3"</code>, the output should be
  <code>bishopAndPawn(bishop, pawn) = true</code>.

  ![](./img.png)

- For <code>bishop = "h1"</code> and <code>pawn = "h3"</code>, the output should be
  <code>bishopAndPawn(bishop, pawn) = false</code>.

  ![](./img_1.png)

**Input/Output**

- **[execution time limit] 4 seconds (js)**

- **[input] string bishop**

  Coordinates of the white bishop in the chess notation.<br>

  _Guaranteed constraints:_<br>
  <code>bishop.length = 2</code>,<br>
  <code>'a' ≤ bishop[0] ≤ 'h'</code>,<br>
  <code>1 ≤ bishop[1] ≤ 8</code>.

  Note (Chess notation): Each square of the chessboard is identified by a unique coordinate pair—a letter and a number. The vertical columns of squares from white's left to the right are labeled 'a' through 'h'. The horizontal rows of squares are numbered 1 to 8 starting from white's side of the board. Thus each square has a unique identification as a string consisting of two characters: the first is the column label, and the second in the row number.

  ![](./img_3.png)

- **[input] string pawn**

  Coordinates of the black pawn in the same notation.<br>

  _Guaranteed constraints:_<br>
  <code>pawn.length = 2</code>,<br>
  <code>'a' ≤ pawn[0] ≤ 'h'</code>,<br>
  <code>1 ≤ pawn[1] ≤ 8</code>.

* **[output] boolean**
  - <code>true</code> if the bishop can capture the pawn, <code>false</code> otherwise.

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
