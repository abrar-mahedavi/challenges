`Medium`	`Codewriting` 	`300`

## Description
------

Given a <code>divisor</code> and a <code>bound</code>, find the largest integer <code>N</code> such that:

* <code>N</code> is divisible by divisor.
* <code>N</code> is less than or equal to bound.
* <code>N</code> is greater than 0.

It is guaranteed that such a number exists.

**Example**

For <code>divisor = 3</code> and <code>bound = 10</code>, the output should be
<code>maxMultiple(divisor, bound) = 9</code>.

The largest integer divisible by 3 and not larger than <code>10</code> is <code>9</code>.

**Input/Output**

* **[execution time limit] 4 seconds (js)**

* **[input] integer divisor**

  _Guaranteed constraints:_<br>
  <code type='math/tex'>2 \leq divisor \leq 10</code>.

* **[input] integer bound**

  _Guaranteed constraints:_<br>
  <code type='math/tex'>5 \leq bound \leq 100</code>.

* **[output] integer**

  * The largest integer not greater than <code>bound</code> that is divisible by <code>divisor</code>.

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
