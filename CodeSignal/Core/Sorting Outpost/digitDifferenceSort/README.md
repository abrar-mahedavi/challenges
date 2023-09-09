`Medium`	`Codewriting` 	`300`

## Description

---

Given an array of integers, sort its elements by the _difference_ of their largest and smallest digits. In the case of a tie, that with the larger index in the array should come first.

**Example**

For <code>a = [152, 23, 7, 887, 243]</code>, the output should be
<code>digitDifferenceSort(a) = [7, 887, 23, 243, 152]</code>.

Here are the _differences_ of all the numbers:

- <code>152</code>: <code>difference = 5 - 1 = 4</code>;
- <code>23</code>: <code>difference = 3 - 2 = 1</code>;
- <code>7</code>: <code>difference = 7 - 7 = 0</code>;
- <code>887</code>: <code>difference = 8 - 7 = 1</code>;
- <code>243</code>: <code>difference = 4 - 2 = 2</code>.

<code>23</code> and <code>887</code> have the same _difference_, but <code>887</code> goes after <code>23</code> in <code>a</code>, so in the sorted array it comes first.

**Input/Output**

- **[execution time limit] 4 seconds (js)**

- **[input] array.integer a**

  An array of integers.

  _Guaranteed constraints:_<br>
  <code>0 ≤ sequence.length ≤ 10<sup>4</sup>,<br>
  <code>1 ≤ sequence[i] ≤ 10<sup>5</sup></code>.

* **[output] array.integer**
  - Array <code>a</code> sorted as described above.
  
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
