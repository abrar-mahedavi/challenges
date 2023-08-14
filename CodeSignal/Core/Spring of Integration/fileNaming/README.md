`Medium`	`Codewriting` 	`300`

## Description

---

You are given an array of desired filenames in the order of their creation. Since two files cannot have equal names, the one which comes later will have an addition to its name in a form of <code>(k)</code>, where <code>k</code> is the smallest positive integer such that the obtained name is not used yet.

Return an array of names that will be given to the files.

**Example**

For <code>names = ["doc", "doc", "image", "doc(1)", "doc"]</code>, the output should be
<code>fileNaming(names) = ["doc", "doc(1)", "image", "doc(1)(1)", "doc(2)"]</code>.

**Input/Output**

- **[execution time limit] 4 seconds (js)**

- **[input] array.string names**

  _Guaranteed constraints:_<br>
  <code>5 ≤ names.length ≤ 1000</code>,<br> <code>1 ≤ names[i].length ≤ 15</code>.

* **[output] array.string**

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
