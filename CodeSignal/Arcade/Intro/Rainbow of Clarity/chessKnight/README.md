`Medium`	`Codewriting` 	`300`

Given a position of a knight on the standard chessboard, find the number of different moves the knight can perform.

The knight can move to a square that is two squares horizontally and one square vertically, or two squares vertically and one square horizontally away from it. The complete move therefore looks like the letter L. Check out the image below to see all valid moves for a knight piece that is placed on one of the central squares.

![knight](knight.jpg)

## Example

- For cell = "a1", the output should be
`chessKnight(cell) = 2`.

    ![knight1](ex_1.jpg)

- For cell = "c2", the output should be
`chessKnight(cell) = 6`.

    ![knight1](ex_2.jpg)

## Input/Output

- [execution time limit] 4 seconds (go)

- [input] string cell

    String consisting of 2 letters - coordinates of the knight on an 8 × 8 chessboard in chess notation.

    Guaranteed constraints: \
    `cell.length = 2`, \
    `'a' ≤ cell[0] ≤ 'h'`, \
    `1 ≤ cell[1] ≤ 8`. \

- [output] integer

## [Go] Syntax Tips

``` go
// Prints help message to the console
// Returns a string
func helloWorld(name string) string {
    fmt.Printf("This prints to the console when you Run Tests");
    return "Hello, " + name;
}
```
