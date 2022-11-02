`Easy`	`Codewriting` 	`300`

Given the string, check if it is a palindrome.

```
A string that doesn't changed when reversed (it reads the same backward and forward).

Examples:

"eye" is a palindrome
"noon" is a palindrome
"decaf faced" is a palindrome
"taco cat" is not a palindrome (backwards it spells "tac ocat")
"racecars" is not a palindrome (backwards it spells "sracecar")
```

## Example

- For `inputString = "aabaa"`, the output should be
`checkPalindrome(inputString) = true`;
- For `inputString = "abac"`, the output should be
`checkPalindrome(inputString) = false`;
- For `inputString = "a"`, the output should be
`checkPalindrome(inputString) = true`.


## Input/Output

- [execution time limit] 4 seconds (go)

- [input] string inputString \
A non-empty string consisting of lowercase characters. \
Guaranteed constraints: \
`1 ≤ inputString.length ≤ 105.`

- [output] boolean \
`true` if `inputString` is a palindrome, `false` otherwise.

## [Go] Syntax Tips

```
// Prints help message to the console
// Returns a string
func helloWorld(name string) string {
    fmt.Printf("This prints to the console when you Run Tests");
    return "Hello, " + name;
}
```