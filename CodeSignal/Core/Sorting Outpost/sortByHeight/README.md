`Medium`	`Codewriting` 	`300`


## Description

---

Some people are standing in a row in a park. There are trees between them which cannot be moved. Your task is to rearrange the people by their heights in a non-descending order without moving the trees. People can be very tall!

**Example**

For <code>a = [-1, 150, 190, 170, -1, -1, 160, 180]</code>, the output should be
<code>sortByHeight(a) = [-1, 150, 160, 170, -1, -1, 180, 190]</code>.

**Input/Output**

- **[execution time limit] 4 seconds (js)**

- **[input] array.integer a**

  If <code>a[i] = -1</code>, then the <code>i<sup>th</sup></code> position is occupied by a tree. Otherwise <code>a[i]</code> is the height of a person standing in the <code>i<sup>th</sup></code> position.

  _Guaranteed constraints:_<br>
  <code>1 ≤ a.length ≤ 1000</code>,<br>
  <code>-1 ≤ a[i] ≤ 1000</code>.

* **[output] array.integer**
  - Sorted array <code>a</code> with all the trees untouched.

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
