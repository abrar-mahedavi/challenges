`Medium`	`Codewriting` 	`300`


## Description

---

Given an integer <code>n</code>, find the _minimal_ <code>k</code> such that

- <code>k = m!</code> (where <code>m! = 1 _ 2 _ ... \* m</code>) for some integer <code>m</code>;
- <code>k >= n</code>.

In other words, find the smallest factorial which is not less than <code>n</code>.

**Example**

For <code>n = 17</code>, the output should be
<code>leastFactorial(n) = 24</code>.

<code>17 < 24 = 4! = 1 _ 2 _ 3 _ 4</code>, while <code>3! = 1 _ 2 \* 3 = 6 < 17</code>)

**Input/Output**

- **[execution time limit] 4 seconds (js)**

- **[input] integer n**

  A positive integer.

  _Guaranteed constraints:_<br>
  <code type='math/tex'>1 \leq n \leq 120</code>.

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
