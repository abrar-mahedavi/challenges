`Medium`	`Codewriting` 	`300`

## Description

---

Given a rectangular matrix of characters, add a border of asterisks(<code>\*</code>) to it.

**Example**

For<br>
<code>
picture = ["abc",<br>
"ded"]
</code>

the output should be <br>

<code>
addBorder(picture) = ["*****",<br>
                      "*abc*",<br>
                      "*ded*",<br>
                      "*****"]
</code>

**Input/Output**

- **[execution time limit] 4 seconds (js)**

- **[input] array.string picture**

  A non-empty array of non-empty equal-length strings.<br>

  _Guaranteed constraints:_<br>
  <code>1 ≤ picture.length ≤ 100</code>,<br> <code>1 ≤ picture[i].length ≤ 100</code>.

- **[output] array.string**
  - The same matrix of characters, framed with a border of asterisks of width <code>1</code>.
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
