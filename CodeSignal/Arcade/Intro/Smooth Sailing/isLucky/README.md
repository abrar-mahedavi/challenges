`Easy`	`Codewriting` 	`300`

Ticket numbers usually consist of an even number of digits. A ticket number is considered lucky if the sum of the first half of the digits is equal to the sum of the second half.

Given a ticket number n, determine if it's lucky or not.

## Example

- For `n = 1230`, the output should be
`isLucky(n) = true`;
- For `n = 239017`, the output should be
`isLucky(n) = false`.

## Input/Output

- [execution time limit] 4 seconds (go)

- [input] integer n

    A ticket number represented as a positive integer with an even number of digits.

    Guaranteed constraints:
    `10 ≤ n < 106`.

- [output] boolean

    true if n is a lucky ticket number, false otherwise.

## [Go] Syntax Tips

``` go
// Prints help message to the console
// Returns a string
func helloWorld(name string) string {
    fmt.Printf("This prints to the console when you Run Tests");
    return "Hello, " + name;
}
```
