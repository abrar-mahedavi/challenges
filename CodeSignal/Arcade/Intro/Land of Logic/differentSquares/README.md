`Medium`	`Codewriting` 	`300`

Given a rectangular matrix containing only digits, calculate the number of different 2 × 2 squares in it.

## Example

For
```
matrix = [[1, 2, 1],
[2, 2, 2],
[2, 2, 2],
[1, 2, 3],
[2, 2, 1]]
```
the output should be
`solution(matrix) = 6`.

Here are all `6` different `2 × 2` squares:

- 1 2 \
  2 2
- 2 1 \
  2 2
- 2 2 \
  2 2
- 2 2 \
  1 2
- 2 2 \
  2 3
- 2 3 \
  2 1

## Input/Output

- [execution time limit] 4 seconds (go)

- [input] array.array.integer matrix

  Guaranteed constraints:
  ```
  1 ≤ matrix.length ≤ 100,
  1 ≤ matrix[i].length ≤ 100,
  0 ≤ matrix[i][j] ≤ 9.
  ```

- [output] integer

  The number of different `2 × 2` squares in matrix.

## [Go] Syntax Tips

``` go
// Prints help message to the console
// Returns a string
func helloWorld(name string) string {
    fmt.Printf("This prints to the console when you Run Tests");
    return "Hello, " + name;
}
```
