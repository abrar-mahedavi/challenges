`Medium`	`Codewriting` 	`300`

## Description

---

Check whether the given string is a subsequence of the plaintext alphabet, given the following definitions:

- subsequence: <code>A</code> is considered a **subsequence** of <code>B</code> if every element from <code>A</code> appears in <code>B</code> in the same order (not necessarily contiguous; there can be other elements in between). In other words, <code>A</code> can be obtained just by deleting elements from B.

  Examples:

  - <code>[3, 7]</code> is a subsequence of <code>[2, 3, 7]</code>
  - <code>[0, 0]</code> is a subsequence of <code>[9, 0, 2, 1, 0]</code>
  - <code>[1, 2,</code> 3] is a subsequence of <code>[1, 2, 3]</code>
  - <code>[8, 7]</code> is not a subsequence of <code>[7, 8] </code>because the elements don't appear in the same order
  - <code>[1, 1]</code> is not a subsequence of <code>[1, 2, 3]</code> because it contains different elements

- plaintext: The **plaintext alphabet** is a string "abcdef...xyz".

**Example**

- For <code>s = "effg"</code>, the output should be
  <code>alphabetSubsequence(s) = false</code>;
- For <code>s = "cdce"</code>, the output should be
  <code>alphabetSubsequence(s) = false</code>;
- For <code>s = "ace"</code>, the output should be
  <code>alphabetSubsequence(s) = true</code>;
- For <code>s = "bxz"</code>, the output should be
  <code>alphabetSubsequence(s) = true</code>.

**Input/Output**

- **[execution time limit] 4 seconds (js)**

- **[input] string s**

  _Guaranteed constraints:_<br>
  <code>2 ≤ s.length ≤ 15</code>.

- **[output] boolean**
  - <code>true</code> if the given string is a _subsequence of the alphabet_, <code>false</code> otherwise.
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
