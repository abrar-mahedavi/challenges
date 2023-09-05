`Medium`	`Codewriting` 	`300`


## Description

---

Given an array of strings, sort them in the order of increasing lengths. If two strings have the same length, their relative order must be the same as in the initial array.

**Example**

For

<code type='preformat'>
inputArray = ["abc",
              "",
              "aaa",
              "a",
              "zz"]
</code>

the output should be

<code type='preformat'>
sortByLength(inputArray) = ["",
                            "a",
                            "zz",
                            "abc",
                            "aaa"]
</code>

**Input/Output**

- **[execution time limit] 4 seconds (js)**

- **[input] array.string inputArray**

  A non-empty array of strings.

  _Guaranteed constraints:_<br>
  <code>3 ≤ inputArray.length ≤ 100</code>,<br>
  <code>0 ≤ inputArray[i].length ≤ 100</code>.

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
