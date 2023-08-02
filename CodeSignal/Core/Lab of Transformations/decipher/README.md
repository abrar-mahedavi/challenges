`Medium`	`Codewriting` 	`300`

## Description

---

Consider the following ciphering algorithm:

- For each character replace it with its code.
- Concatenate all of the obtained numbers.

Given a ciphered string, return the initial one if it is known that it consists only of lowercase letters.

Note: here the _character's code_ means its decimal ASCII code, the numerical representation of a character used by most modern programming languages.

**Example**

For <code>cipher = "10197115121"</code>, the output should be
<code>decipher(cipher) = "easy"</code>.

Explanation: <code>charCode('e') = 101</code>, <code>charCode('a') = 97</code>, <code>charCode('s') = 115</code> and <code>charCode('y') = 121</code>.

**Input/Output**

- **[execution time limit] 4 seconds (js)**

- **[input] string cipher**

  A non-empty string which is guaranteed to be a cipher for some other string of lowercase letters.

  _Guaranteed constraints:_<br>
  <code>2 ≤ cipher.length ≤ 100</code>.

- **[output] string**

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
