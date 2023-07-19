`Medium`	`Codewriting` 	`300`

## Description

---

<code>N</code> candles are placed in a row, some of them are initially lit. For each candle from the <code>1st</code> to the <code>Nth</code> the following algorithm is applied: if the observed candle is lit then states of this candle and all candles before it are changed to the opposite. Which candles will remain lit after applying the algorithm to all candles in the order they are placed in the line?

**Example**

- For <code>a = [1, 1, 1, 1, 1]</code>, the output should be
  <code>switchLights(a) = [0, 1, 0, 1, 0]</code>.<br>
  Check out the image below for better understanding:<br>

  ![](./images/example.png)

- For <code>a = [0, 0]</code>, the output should be
  <code>switchLights(a) = [0, 0]</code>.

  The candles are not initially lit, so their states are not altered by the algorithm.

</code>

**Input/Output**

- **[execution time limit] 4 seconds (js)**

- **[input] array.integer a**

  Initial situation - array of zeros and ones of length N, 1 means that the corresponding candle is lit.<br>

  _Guaranteed constraints:_<br>
  <code>2 ≤ a.length ≤ 5000</code>.

- **[output] array.integer**
  - Situation after applying the algorithm - array in the same format as input with the same length.
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
