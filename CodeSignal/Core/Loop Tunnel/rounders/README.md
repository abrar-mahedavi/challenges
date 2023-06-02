`Medium`	`Codewriting` 	`300`


## Description

---

We want to turn the given integer into a number that has only one non-zero digit using a tail rounding approach. This means that at each step we take the last non <code>0</code> digit of the number and round it to <code>0</code> or to <code>10</code>. If it's less than <code>5</code> we round it to <code>0</code> if it's larger than or equal to <code>5</code> we round it to <code>10</code> (rounding to <code>10</code> means increasing the next significant digit by <code>1</code>). The process stops immediately once there is only one non-zero digit left.

**Example**

- For <code>n = 15</code>, the output should be
  <code>rounders(n) = 20</code>;

- For <code>n = 1234</code>, the output should be
  <code>rounders(n) = 1000</code>.

  <code>1234 -> 1230 -> 1200 -> 1000</code>.

- For <code>n = 1445</code>, the output should be
  <code>rounders(n) = 2000</code>.

  <code>1445 -> 1450 -> 1500 -> 2000</code>.

**Input/Output**

- **[execution time limit] 4 seconds (js)**

- **[input] integer n**

  A positive integer.

  _Guaranteed constraints:_<br>
  <code type='math/tex'>1 \leq value \leq 10^8</code>.

- **[output] integer**

  - The rounded number.

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
