`Easy`	`Codewriting` 	`300`

Write a function that reverses characters in (possibly nested) parentheses in the input string.

Input strings will always be well-formed with matching ()s.

## Example

- For `inputString = "(bar)"`, the output should be
`reverseInParentheses(inputString) = "rab"`;

- For `inputString = "foo(bar)baz"`, the output should be
`reverseInParentheses(inputString) = "foorabbaz"`;

- For `inputString = "foo(bar)baz(blim)"`, the output should be
`reverseInParentheses(inputString) = "foorabbazmilb"`;

- For `inputString = "foo(bar(baz))blim"`, the output should be
`reverseInParentheses(inputString) = "foobazrabblim"`.
    Because `"foo(bar(baz))blim"` becomes `"foo(barzab)blim"` and then `"foobazrabblim"`.

## Input/Output

- [execution time limit] 4 seconds (go)

- [input] array.array.integer matrix \
A 2-dimensional array of integers representing the cost of each room in the building. A value of 0 indicates that the room is haunted. \
Guaranteed constraints: \
`1 ≤ matrix.length ≤ 5`, \
`1 ≤ matrix[i].length ≤ 5`, \
`0 ≤ matrix[i][j] ≤ 10`.

- [output] integer \
The total price of all the rooms that are suitable for the CodeBots to live in.

## [Go] Syntax Tips

``` go
// Prints help message to the console
// Returns a string
func helloWorld(name string) string {
    fmt.Printf("This prints to the console when you Run Tests");
    return "Hello, " + name;
}
```
