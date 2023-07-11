`Medium`	`Codewriting` 	`300`

## Description

---

Let's call two integers <code>A</code> and <code>B</code> _friends_ if each integer from the array <code>divisors</code> is either a divisor of both <code>A</code> and <code>B</code> or neither <code>A</code> nor <code>B</code>. If two integers are _friends_, they are said to be in the same _clan_. How many clans are the integers from <code>1</code> to <code>k</code>, inclusive, broken into?

**Example**

For <code>divisors = [2, 3]</code> and <code>k = 6</code>, the output should be
<code>numberOfClans(divisors, k) = 4</code>.

The numbers <code>1</code> and <code>5</code> are friends and form a clan, <code>2</code> and <code>4</code> are friends and form a clan, and <code>3</code> and <code>6</code> do not have friends and each is a clan by itself. So the numbers <code>1</code> through <code>6</code> are broken into <code>4</code> clans.

**Input/Output**

- **[execution time limit] 4 seconds (js)**

- **[input] array.integer divisors**

  A non-empty array of positive integers.<br>

  _Guaranteed constraints:_<br>
  <code>2 ≤ divisors.length < 10</code>,<br> <code>1 ≤ divisors[i] ≤ 10</code>.

- **[input] integer k**

  A positive integer.<br>

  _Guaranteed constraints:_<br>
  <code>5 ≤ k ≤ 10</code>.

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
