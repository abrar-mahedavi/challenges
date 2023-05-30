`Medium`	`Codewriting` 	`300`

## Description

---

A little boy is studying arithmetics. He has just learned how to add two integers, written one below another, column by column. But he always forgets about the important part - carrying.

Given two integers, your task is to find the result which the little boy will get.

_Note: The boy had learned [from](https://www.mathsisfun.com/numbers/addition-column.html) this site, so feel free to check it out too if you are not familiar with column addition_.

**Example**

For <code>param1 = 456</code> and <code>param2 = 1734</code>, the output should be
<code>additionWithoutCarrying(param1, param2) = 1180</code>.

<code>
&nbsp;&nbsp;&nbsp;456<br>
&nbsp;&nbsp;1734<br>
+ ____<br>
&nbsp;&nbsp;1180
</code>

The boy performs the following operations from right to left:

- <code>6 + 4 = 10</code> but he forgets about carrying the <code>1</code> and just writes down the <code>0</code> in the last column
- <code>5 + 3 = 8</code>
- <code>4 + 7 = 11</code> but he forgets about the leading <code>1</code> and just writes down <code>1</code> under <code>4</code> and <code>7</code>.
- There is no digit in the first number corresponding to the leading digit of the second one, so the boy imagines that <code>0</code> is written before <code>456</code>. Thus, he gets <code>0 + 1 = 1</code>.

**Input/Output**

- **[execution time limit] 4 seconds (js)**

- **[input] integer param1**

  A non-negative integer.<br>

  _Guaranteed constraints:_<br>
  <code type='math/tex'>0 \leq param1 \le 10^5</code>.

- **[input] integer param1**

  A non-negative integer.<br>

  _Guaranteed constraints:_<br>
  <code type='math/tex'>0 \leq param2 \le 6 · 10^4</code>.

- **[output] integer**
  - The result that the little boy will get by using column addition without carrying.

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
