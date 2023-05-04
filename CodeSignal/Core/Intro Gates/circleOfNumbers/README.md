`Medium`	`Codewriting` 	`300`

## Description
------

Consider integer numbers from <code>0</code> to <code>n - 1</code> written down along the circle in such a way that the distance between any two neighboring numbers is equal (note that <code>0</code> and <code>n - 1</code> are neighboring, too).

Given <code>n</code> and <code>firstNumber</code>, find the number which is written in the radially opposite position to <code>firstNumber</code>.


**Example**

For <code>n = 10</code> and <code>firstNumber = 2</code>, the output should be
<code>circleOfNumbers(n, firstNumber) = 7</code>.

![](./images/example.png)

**Input/Output**

* **[execution time limit] 4 seconds (js)**

* **[input] integer n**

  A positive **even** integer.

  _Guaranteed constraints:_<br>
  <code type='math/tex'>4 \leq divisor \leq 20</code>.

* **[input] integer firstNumber**

  _Guaranteed constraints:_<br>
  <code type='math/tex'>0 \leq firstNumber \leq n - 1</code>.

* **[output] integer**

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
