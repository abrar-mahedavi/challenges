`Easy`	`Codewriting` 	`300`

Given two cells on the standard chess board, determine whether they have the same color or not.

## Example

- For `cell1 = "A1"` and `cell2 = "C3"`, the output should be `chessBoardCellColor(cell1, cell2) = true`.

    ![exmaple1.png](example1.png)

- For `cell1 = "A1"`and `cell2 = "H3"`, the output should be `chessBoardCellColor(cell1, cell2) = false`.

    ![exmaple2.png](example2.png)

## Input/Output

- [execution time limit] 4 seconds (go)

- [input] string cell1

    Guaranteed constraints: \
    `cell1.length = 2`, \
    `'A' ≤ cell1[0] ≤ 'H'`, \
    `1 ≤ cell1[1] ≤ 8`.

- [input] string cell2

    Guaranteed constraints: \
    `cell2.length = 2`, \
    `'A' ≤ cell2[0] ≤ 'H'`, \
    `1 ≤ cell2[1] ≤ 8`.

- [output] boolean

    `true` if both cells have the same color, `false` otherwise.

## [Go] Syntax Tips

``` go
// Prints help message to the console
// Returns a string
func helloWorld(name string) string {
    fmt.Printf("This prints to the console when you Run Tests");
    return "Hello, " + name;
}
```
