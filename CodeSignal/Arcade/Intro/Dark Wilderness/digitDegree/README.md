`Medium`	`Codewriting` 	`300`

Let's define digit degree of some positive integer as the number of times we need to replace this number with the sum of its digits until we get to a one digit number.

Given an integer, find its digit degree.

## Example

- For `n = 5`, the output should be
`digitDegree(n) = 0`;

- For `n = 100`, the output should be
`digitDegree(n) = 1`.
`1 + 0 + 0 = 1`.

- For `n = 91`, the output should be
`digitDegree(n) = 2`.
`9 + 1 = 10 -> 1 + 0 = 1`.

## Input/Output

- [execution time limit] 4 seconds (go)

- [input] integer n

    Guaranteed constraints: \
    `5 ≤ n ≤ 109`.

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
