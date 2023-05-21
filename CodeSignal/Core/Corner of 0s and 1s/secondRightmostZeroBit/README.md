`Medium`	`Codewriting` 	`300`

## Description

---

_Implement the missing code, denoted by ellipses. You may not modify the pre-existing code._

Presented with the integer <code>n</code>, find the 0-based position of the second rightmost zero bit in its binary representation (it is guaranteed that such a bit exists), counting from right to left.

Return the value of <code>2<sup>position_of_the_found_bit</sup></code>.

**Example**

For <code>n = 37</code>, the output should be
<code>secondRightmostZeroBit(n) = 8</code>.

<code>37<sub>10</sub> = 10<b>0</b>101<sub>2</sub></code>. The second rightmost zero bit is at position <code>3</code> (0-based) from the right in the binary representation of <code>n</code>.
Thus, the answer is <code>2<sup>3</sup> = 8</code>.

**Input/Output**

- **[execution time limit] 4 seconds (js)**

- **[input] integer n**

  _Guaranteed constraints:_<br>
  <code type='math/tex'>4 \leq n \leq 2^{30}</code>.

- **[output] integer**

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
