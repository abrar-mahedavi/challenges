`Medium`	`Codewriting` 	`300`

## Description

---

Remove a part of a given array between given 0-based indexes <code>l</code> and <code>r</code> (inclusive).

**Example**

For <code>inputArray = [2, 3, 2, 3, 4, 5]</code>, <code>l = 2</code>, and <code>r = 4</code>, the output should be
<code>removeArrayPart(inputArray, l, r) = [2, 3, 5]</code>.

**Input/Output**

- **[execution time limit] 4 seconds (js)**

- **[input] array.integer inputArray**

  _Guaranteed constraints:_<br>
  <code>2 ≤ inputArray.length ≤ 10<sup>4</sup></code>,<br>
  <code>-10<sup>5</sup> ≤ inputArray[i] ≤ 10<sup>5</sup></code>.

- **[input] integer l**

  Left index of the part to be removed (<code>0</code>-based).<br>

  _Guaranteed constraints:_<br>
  <code>0 ≤ l ≤ r</code>.

- **[input] integer r**

  Right index of the part to be removed (<code>0</code>-based).<br>

  _Guaranteed constraints:_<br>
  <code>l ≤ r < inputArray.length</code>.

- **[output] array.integer**

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
