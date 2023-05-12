`Medium`	`Codewriting` 	`300`


## Description
------

Consider an arithmetic expression of the form <code>a#b=c</code>. Check whether it is possible to replace <code>#</code> with one of the four signs: <code>+</code>, <code>-</code>, <code>*</code> or <code>/</code> to obtain a correct expression.


**Example**

* For <code>a = 2</code>, <code>b = 3</code>, and <code>c = 5</code>, the output should be
  <code>arithmeticExpression(a, b, c) = true</code>.

We can replace <code>#</code> with a <code>+</code> to obtain <code>2 + 3 = 5</code>, so the answer is <code>true</code>.

* For <code>a = 8</code>, <code>b = 2</code>, and <code>c = 4</code>, the output should be
  <code>arithmeticExpression(a, b, c) = true</code>.

We can replace <code>#</code> with a <code>/</code> to obtain <code>8 / 2 = 4</code>, so the answer is <code>true</code>.

* For <code>a = 8</code>, <code>b = 3</code>, and <code>c = 2</code>, the output should be
  <code>arithmeticExpression(a, b, c) = false</code>.
  * <code>8 + 3 = 11 ≠ 2</code>;
  * <code>8 - 3 = 5 ≠ 2</code>;
  * <code>8 * 3 = 24 ≠ 2</code>;
  * <code>8 / 3 = 2.(6) ≠ 2</code>.

  So the answer is <code>false</code>.


**Input/Output**

* **[execution time limit] 4 seconds (js)**

* **[input] integer a**

  _Guaranteed constraints:_<br>
  <code type='math/tex'>1 \leq a \leq 20</code>.

* **[input] integer b**

  _Guaranteed constraints:_<br>
  <code type='math/tex'>1 \leq b \leq 20</code>.

* **[input] integer c**

  _Guaranteed constraints:_<br>
  <code type='math/tex'>0 \leq c \leq 20</code>.

* **[output] boolean**

  * <code>true</code> if the desired replacement is possible, <code>false</code> otherwise.

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
