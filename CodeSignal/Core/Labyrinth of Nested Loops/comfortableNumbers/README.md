`Medium`	`Codewriting` 	`300`

## Description

---

Let's say that number <code>a</code> feels comfortable with number <code>b</code> if <code>a ≠ b</code> and <code>b</code> lies in the segment <code>[a - s(a), a + s(a)]</code>, where <code>s(x)</code> is the sum of <code>x</code>'s digits.

How many pairs <code>(a, b)</code> are there, such that <code>a < b</code>, both <code>a</code> and <code>b</code> lie on the segment <code>[l, r]</code>, and each number feels comfortable with the other (so a feels comfortable with <code>b</code> and <code>b</code> feels comfortable with <code>a</code>)?

**Example**

For <code>l = 10</code> and <code>r = 12</code>, the output should be
<code>comfortableNumbers(l, r) = 2</code>.

Here are all values of <code>s(x)</code> to consider:

- <code>s(10) = 1</code>, so <code>10</code> is comfortable with <code>9</code> and <code>11</code>;
- <code>s(11) = 2</code>, so <code>11</code> is comfortable with <code>9</code>, <code>10</code>,<code>12</code> and <code>13</code>;
- <code>s(12) = 3</code>, so <code>12</code> is comfortable with <code>9</code>, <code>10</code>, <code>11</code>, <code>13</code>, <code>14</code> and <code>15</code>.

Thus, there are <code>2</code> pairs of numbers comfortable with each other within the segment <code>[10; 12]</code>: <code>(10, 11)</code> and <code>(11, 12)</code>.

**Input/Output**

- **[execution time limit] 4 seconds (js)**

- **[input] integer l**

  _Guaranteed constraints:_<br>
  <code>1 ≤ l ≤ r ≤ 1000</code>.

- **[input] integer r**

  _Guaranteed constraints:_<br>
  <code>1 ≤ l ≤ r ≤ 1000</code>.

- **[output] integer**
  - The number of pairs satisfying all the above conditions.

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
