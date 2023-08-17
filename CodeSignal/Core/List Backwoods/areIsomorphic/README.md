`Medium`	`Codewriting` 	`300`

## Description

---

Two two-dimensional arrays are _isomorphic_ if they have the same number of rows and each pair of respective rows contains the same number of elements.

Given two two-dimensional arrays, check if they are isomorphic.

**Example**

- For
  <code type='preformat'>
  array1 = [[1, 1, 1],
  [0, 0]]
  </code>
  and

  <code type='preformat'>
  array2 = [[2, 1, 1],
            [2, 1]]
  </code>

  the output should be

  <code>areIsomorphic(array1, array2) = true;</code>

- For
  <code type='preformat'>
  array1 = [[2],
  []]
  </code>

  and
  <code type='preformat'>
  array2 = [[2]]
  </code>

  the output should be

  <code>areIsomorphic(array1, array2) = false.</code>

**Input/Output**

- **[execution time limit] 4 seconds (js)**

- **[input] array.array.integer array1**

  _Guaranteed constraints:_<br>
  <code>1 ≤ array1.length ≤ 5</code>,<br> <code>0 ≤ array1[i].length ≤ 5</code>,<br> <code>0 ≤ array1[i][j] ≤ 50</code>.

* **[input] array.array.integer array2**

  _Guaranteed constraints:_<br>
  <code>1 ≤ array2.length ≤ 5</code>,<br> <code>0 ≤ array2[i].length ≤ 5</code>,<br> <code>0 ≤ array2[i][j] ≤ 50</code>.

* **[output] boolean**

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
